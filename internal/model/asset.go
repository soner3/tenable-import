package model

type TenableImportPayload struct {
	Source string  `json:"source"`
	Assets []Asset `json:"assets"`
}

type Asset struct {
	FQDNs           []string `json:"fqdn,omitempty"`
	IPv4s           []string `json:"ipv4,omitempty"`
	MACAddresses    []string `json:"mac_address,omitempty"`
	NetBIOS         string   `json:"netbios_name,omitempty"`
	OperatingSystem []string `json:"operating_system,omitempty"`
	BIOSUUID        string   `json:"bios_uuid,omitempty"`
}

type TenableExportResponseAsset struct {
	ID      string `json:"id"`
	Network struct {
		IPv4s             []string `json:"ipv4s,omitempty"`
		FQDNs             []string `json:"fqdns,omitempty"`
		NetworkInterfaces []struct {
			IPv4s []string `json:"ipv4s,omitempty"`
			FQDNs []string `json:"fqdns,omitempty"`
		} `json:"network_interfaces,omitempty"`
	} `json:"network"`
}
