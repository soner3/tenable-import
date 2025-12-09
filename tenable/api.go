package tenable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const BaseURL = "https://cloud.tenable.com"
const MaxRetrys = 5
const DefaultWaitSeconds = 5
const StatusCheckInterval = 10 * time.Second

// ExportAssetsV2 startet einen Asset-Export mit chunk_size 5000,
// wartet bis der Export fertig ist und lädt alle Chunks herunter
func (c *TenableClient) ExportAssetsV2(filters *AssetExportFilters) ([]*Asset, error) {
	exportReq := AssetExportRequest{
		ChunkSize: 5000,
		Filters:   filters,
	}

	var exportResp AssetExportResponse
	_, err := c.CallAPI(http.MethodPost, "/assets/v2/export", exportReq, &exportResp)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Starten des Asset-Exports: %w", err)
	}

	exportUUID := exportResp.ExportUUID
	if exportUUID == "" {
		return nil, fmt.Errorf("export UUID ist leer")
	}

	c.App.DebugLogger.Println("Warte auf Abschluss des Exports...")

	var statusResp AssetExportStatusResponse
	statusCheckCount := 0
	for {
		statusCheckCount++
		_, err = c.CallAPI(http.MethodGet, fmt.Sprintf("/assets/export/%s/status", exportUUID), nil, &statusResp)
		if err != nil {
			c.App.ErrorLogger.Printf("Fehler beim Abrufen des Export-Status: %v", err)
			return nil, fmt.Errorf("fehler beim Abrufen des Export-Status: %w", err)
		}

		c.App.DebugLogger.Printf("Status-Check #%d: %s (Chunks: %d)", statusCheckCount, statusResp.Status, len(statusResp.ChunksAvailable))

		switch statusResp.Status {
		case "FINISHED":
			c.App.InfoLogger.Printf("Export abgeschlossen - %d Chunks verfügbar", len(statusResp.ChunksAvailable))
			allAssets := make([]*Asset, 0)

			for i, chunkID := range statusResp.ChunksAvailable {
				c.App.DebugLogger.Printf("Lade Chunk %d/%d (ID: %d)", i+1, len(statusResp.ChunksAvailable), chunkID)
				var chunkAssets []*Asset
				_, err = c.CallAPI(
					http.MethodGet,
					fmt.Sprintf("/assets/export/%s/chunks/%d", exportUUID, chunkID),
					nil,
					&chunkAssets,
				)
				if err != nil {
					c.App.ErrorLogger.Printf("Fehler beim Herunterladen von Chunk %d: %v", chunkID, err)
					return nil, fmt.Errorf("fehler beim Herunterladen von Chunk %d: %w", chunkID, err)
				}

				c.App.InfoLogger.Printf("Chunk %d/%d geladen: %d Assets", i+1, len(statusResp.ChunksAvailable), len(chunkAssets))
				allAssets = append(allAssets, chunkAssets...)
			}

			c.App.InfoLogger.Printf("Export erfolgreich: %d Assets geladen", len(allAssets))
			return allAssets, nil
		case "QUEUED", "PROCESSING":
			c.App.TraceLogger.Printf("Export läuft noch, warte %v...", StatusCheckInterval)
			time.Sleep(StatusCheckInterval)
		case "CANCELLED":
			c.App.WarnLogger.Println("Export wurde abgebrochen")
			return nil, fmt.Errorf("export wurde abgebrochen")
		case "ERROR":
			c.App.ErrorLogger.Println("Export ist fehlgeschlagen")
			return nil, fmt.Errorf("fehler beim Export")
		default:
			c.App.ErrorLogger.Printf("Unbekannter Export-Status: %s", statusResp.Status)
			return nil, fmt.Errorf("unbekannter Export-Status: %s", statusResp.Status)
		}
	}

}

func (c *TenableClient) CallAPI(method string, endpoint string, reqBody, resBody any) (int, error) {
	c.App.TraceLogger.Printf("API Call: %s %s", method, endpoint)

	b, err := json.Marshal(reqBody)
	if err != nil {
		c.App.ErrorLogger.Printf("JSON Marshal Fehler: %v", err)
		return 0, err
	}
	req, err := http.NewRequest(method, BaseURL+endpoint, bytes.NewReader(b))
	if err != nil {
		c.App.ErrorLogger.Printf("HTTP Request Fehler: %v", err)
		return 0, err
	}

	req.Header.Add("X-ApiKeys", c.App.TenableAPIKey)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.App.ErrorLogger.Printf("HTTP Do Fehler: %v", err)
		return 0, err
	}

	fulfilled := false
	retryCount := 0

	for !fulfilled {
		switch res.StatusCode {
		case http.StatusOK, http.StatusCreated, http.StatusAccepted:
			fulfilled = true
		case http.StatusBadRequest:
			c.App.ErrorLogger.Println("400 Bad Request: Ungültige Parameter")
			return http.StatusBadRequest, fmt.Errorf("400 Bad Request: üngültige Parameter oder Formatierung der Anfrage")
		case http.StatusUnauthorized:
			c.App.ErrorLogger.Println("401 Unauthorized: Ungültiger API-Schlüssel")
			return http.StatusUnauthorized, fmt.Errorf("401 Unauthorized: ungültiger oder fehlender API-Schlüssel")
		case http.StatusForbidden:
			c.App.ErrorLogger.Println("403 Forbidden: Unzureichende Berechtigungen")
			return http.StatusForbidden, fmt.Errorf("403 Forbidden: Zugriff verweigert, unzureichende Berechtigungen")
		case http.StatusNotFound:
			c.App.WarnLogger.Printf("404 Not Found: %s", endpoint)
			return http.StatusNotFound, fmt.Errorf("404 Not Found: Ressource nicht gefunden")
		case http.StatusConflict:
			c.App.WarnLogger.Println("409 Conflict: Ressourcen-Konflikt")
			return http.StatusConflict, fmt.Errorf("409 Conflict: Konflikt mit dem aktuellen Zustand der Ressource")
		case http.StatusTooManyRequests:
			retryCount += 1
			if retryCount >= MaxRetrys {
				c.App.ErrorLogger.Printf("Maximale Retry-Anzahl erreicht (%d)", MaxRetrys)
				return 0, fmt.Errorf("maximale Anzahl an Wiederholungen erreicht (%d)", MaxRetrys)
			}
			retryAfter := res.Header.Get("retry-after")
			var waitSeconds int
			if retryAfter != "" {
				_, err = fmt.Sscanf(retryAfter, "%d", &waitSeconds)
				if err != nil {
					return 0, err
				}
			} else {
				waitSeconds = DefaultWaitSeconds
			}
			c.App.WarnLogger.Printf("429 Too Many Requests - Retry #%d in %ds", retryCount, waitSeconds)
			retryCount++
			if retryCount >= MaxRetrys {
				return 0, fmt.Errorf("maximale Anzahl an Wiederholungen erreicht (%d)", MaxRetrys)
			}
			time.Sleep(time.Duration(waitSeconds) * time.Second)
		case http.StatusInternalServerError:
			c.App.ErrorLogger.Println("500 Internal Server Error")
			return 0, fmt.Errorf("500 Internal Server Error: Serverfehler, bitte später erneut versuchen")
		default:
			c.App.ErrorLogger.Printf("Unerwarteter Status Code: %d", res.StatusCode)
			return 0, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(resBody); err != nil {
		c.App.ErrorLogger.Printf("JSON Decode Fehler: %v", err)
		return 0, err
	}

	c.App.TraceLogger.Printf("API Call erfolgreich: %s %s", method, endpoint)
	return 0, nil
}
