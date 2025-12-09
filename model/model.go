package model

type Asset struct {
	FQDNs           []string `json:"fqdn,omitempty"`
	IPv4s           []string `json:"ipv4,omitempty"`
	MACAddresses    []string `json:"mac_address,omitempty"`
	NetBIOS         string   `json:"netbios_name,omitempty"`
	OperatingSystem []string `json:"operating_system,omitempty"`
	BIOSUUID        string   `json:"bios_uuid,omitempty"`
}
