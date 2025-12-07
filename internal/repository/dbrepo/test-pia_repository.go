package dbrepo

import "github.com/soner3/tenable-import/internal/model"

// GetAllOpenAssets ruft alle Test-PIA-Assets ab
func (r *testPiaRepository) GetAllOpenAssets() ([]*model.Asset, error) {
	return []*model.Asset{
		{
			FQDNs:           []string{"server-existing.lan"},
			IPv4s:           []string{"192.168.1.10"},
			MACAddresses:    []string{"00:11:22:33:44:AA"},
			NetBIOS:         "SRV-EXISTING",
			OperatingSystem: []string{"Linux Ubuntu 20.04"},
			BIOSUUID:        "uuid-bios-1",
		},
		{
			FQDNs:           []string{"server-new.lan"},
			IPv4s:           []string{"192.168.1.20"},
			MACAddresses:    []string{"00:11:22:33:44:BB"},
			NetBIOS:         "SRV-NEW",
			OperatingSystem: []string{"Windows Server 2019"},
			BIOSUUID:        "uuid-bios-2",
		},
	}, nil
}

// GetAllClosedAssets ruft alle geschlossenen Test-PIA-Assets ab
func (r *testPiaRepository) GetAllClosedAssets() ([]*model.Asset, error) {
	return []*model.Asset{
		{
			FQDNs:           []string{"server-existing.lan"},
			IPv4s:           []string{"192.168.1.10"},
			MACAddresses:    []string{"00:11:22:33:44:AA"},
			NetBIOS:         "SRV-EXISTING",
			OperatingSystem: []string{"Linux Ubuntu 20.04"},
			BIOSUUID:        "uuid-bios-1",
		},
		{
			FQDNs:           []string{"server-new.lan"},
			IPv4s:           []string{"192.168.1.20"},
			MACAddresses:    []string{"00:11:22:33:44:BB"},
			NetBIOS:         "SRV-NEW",
			OperatingSystem: []string{"Windows Server 2019"},
			BIOSUUID:        "uuid-bios-2",
		},
	}, nil
}
