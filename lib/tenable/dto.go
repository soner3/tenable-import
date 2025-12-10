package tenable

// AssetExportRequest repräsentiert den Body für POST /assets/v2/export
type AssetExportRequest struct {
	// Pflichtfeld: Empfohlen 5000, max 10000
	ChunkSize int `json:"chunk_size"`
	// Optional: Standardmäßig false
	IncludeOpenPorts bool `json:"include_open_ports,omitempty"`
	// Optional: Limitiert ChunkSize auf 1000
	IncludeResourceTags bool `json:"include_resource_tags,omitempty"`
	// Optional: Filterobjekt
	Filters *AssetExportFilters `json:"filters,omitempty"`
}

// AssetExportFilters definiert die Filterkriterien
type AssetExportFilters struct {
	Since                 int64 `json:"since,omitempty"`
	CreatedAt             int64 `json:"created_at,omitempty"`
	UpdatedAt             int64 `json:"updated_at,omitempty"`
	TerminatedAt          int64 `json:"terminated_at,omitempty"`
	DeletedAt             int64 `json:"deleted_at,omitempty"`
	FirstScanTime         int64 `json:"first_scan_time,omitempty"`
	LastAuthenticatedScan int64 `json:"last_authenticated_scan_time,omitempty"`
	LastAssessed          int64 `json:"last_assessed,omitempty"`

	Sources []string `json:"sources,omitempty"`
	Types   []string `json:"types,omitempty"`

	IsTerminated     *bool `json:"is_terminated,omitempty"`
	IsDeleted        *bool `json:"is_deleted,omitempty"`
	IsLicensed       *bool `json:"is_licensed,omitempty"`
	HasPluginResults *bool `json:"has_plugin_results,omitempty"`
	ServiceNowSysID  *bool `json:"servicenow_sysid,omitempty"`

	NetworkID  string `json:"network_id,omitempty"`
	LastScanID string `json:"last_scan_id,omitempty"`
}

// AssetExportResponse ist die Antwort der API (200 OK)
type AssetExportResponse struct {
	ExportUUID string `json:"export_uuid"`
}

// AssetExportStatusResponse ist die Antwort für GET /assets/export/{export_uuid}/status
type AssetExportStatusResponse struct {
	// Status: QUEUED, PROCESSING, FINISHED, CANCELLED, ERROR
	Status          string `json:"status"`
	ChunksAvailable []int  `json:"chunks_available"`
}

// Asset repräsentiert ein Asset aus der v2 API (vereinfacht)
type Asset struct {
	ID                string           `json:"id"`
	HasAgent          bool             `json:"has_agent"`
	HasPluginResults  bool             `json:"has_plugin_results"`
	IsLicensed        bool             `json:"is_licensed"`
	Types             []string         `json:"types,omitempty"`
	OperatingSystems  []string         `json:"operating_systems,omitempty"`
	SystemTypes       []string         `json:"system_types,omitempty"`
	InstalledSoftware []string         `json:"installed_software,omitempty"`
	Sources           []AssetSource    `json:"sources,omitempty"`
	Tags              []AssetTag       `json:"tags,omitempty"`
	Network           *AssetNetwork    `json:"network,omitempty"`
	Timestamps        *AssetTimestamps `json:"timestamps,omitempty"`
	Scan              *AssetScan       `json:"scan,omitempty"`
	Cloud             *AssetCloud      `json:"cloud,omitempty"`
	Ratings           *AssetRatings    `json:"ratings,omitempty"`
}

type AssetSource struct {
	Name      string `json:"name"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}

type AssetTag struct {
	UUID    string `json:"uuid"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	AddedBy string `json:"added_by,omitempty"`
	AddedAt string `json:"added_at,omitempty"`
}

type AssetNetwork struct {
	NetworkID       string             `json:"network_id,omitempty"`
	NetworkName     string             `json:"network_name,omitempty"`
	IPv4s           []string           `json:"ipv4s,omitempty"`
	IPv6s           []string           `json:"ipv6s,omitempty"`
	FQDNs           []string           `json:"fqdns,omitempty"`
	MACAddresses    []string           `json:"mac_addresses,omitempty"`
	Hostnames       []string           `json:"hostnames,omitempty"`
	NetBIOSNames    []string           `json:"netbios_names,omitempty"`
	SSHFingerprints []string           `json:"ssh_fingerprints,omitempty"`
	Interfaces      []NetworkInterface `json:"network_interfaces,omitempty"`
	OpenPorts       []OpenPort         `json:"open_ports,omitempty"`
}

type NetworkInterface struct {
	Name         string   `json:"name"`
	MACAddresses []string `json:"mac_addresses,omitempty"`
	IPv4s        []string `json:"ipv4s,omitempty"`
	IPv6s        []string `json:"ipv6s,omitempty"`
	FQDNs        []string `json:"fqdns,omitempty"`
	Virtual      *bool    `json:"virtual,omitempty"`
	Aliased      *bool    `json:"aliased,omitempty"`
}

type OpenPort struct {
	Port         int      `json:"port"`
	Protocol     string   `json:"protocol"`
	ServiceNames []string `json:"service_names,omitempty"`
	FirstSeen    string   `json:"first_seen,omitempty"`
	LastSeen     string   `json:"last_seen,omitempty"`
}

type AssetTimestamps struct {
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	DeletedAt    string `json:"deleted_at,omitempty"`
	TerminatedAt string `json:"terminated_at,omitempty"`
	FirstSeen    string `json:"first_seen,omitempty"`
	LastSeen     string `json:"last_seen,omitempty"`
}

type AssetScan struct {
	FirstScanTime                 string `json:"first_scan_time,omitempty"`
	LastScanTime                  string `json:"last_scan_time,omitempty"`
	LastAuthenticatedScanDate     string `json:"last_authenticated_scan_date,omitempty"`
	LastLicensedScanDate          string `json:"last_licensed_scan_date,omitempty"`
	LastScanID                    string `json:"last_scan_id,omitempty"`
	LastScheduleID                string `json:"last_schedule_id,omitempty"`
	LastAuthenticationAttemptDate string `json:"last_authentication_attempt_date,omitempty"`
	LastAuthenticationSuccessDate string `json:"last_authentication_success_date,omitempty"`
	LastAuthenticationScanStatus  string `json:"last_authentication_scan_status,omitempty"`
	LastScanTarget                string `json:"last_scan_target,omitempty"`
}

type AssetCloud struct {
	AWS   *CloudAWS   `json:"aws,omitempty"`
	Azure *CloudAzure `json:"azure,omitempty"`
	GCP   *CloudGCP   `json:"gcp,omitempty"`
}

type CloudAWS struct {
	EC2InstanceAMIID     string `json:"ec2_instance_ami_id,omitempty"`
	EC2InstanceID        string `json:"ec2_instance_id,omitempty"`
	OwnerID              string `json:"owner_id,omitempty"`
	AvailabilityZone     string `json:"availability_zone,omitempty"`
	Region               string `json:"region,omitempty"`
	VPCID                string `json:"vpc_id,omitempty"`
	EC2InstanceGroupName string `json:"ec2_instance_group_name,omitempty"`
	EC2InstanceStateName string `json:"ec2_instance_state_name,omitempty"`
	EC2InstanceType      string `json:"ec2_instance_type,omitempty"`
	SubnetID             string `json:"subnet_id,omitempty"`
	EC2ProductCode       string `json:"ec2_product_code,omitempty"`
	EC2Name              string `json:"ec2_name,omitempty"`
}

type CloudAzure struct {
	VMID       string `json:"vm_id,omitempty"`
	ResourceID string `json:"resource_id,omitempty"`
}

type CloudGCP struct {
	ProjectID  string `json:"project_id,omitempty"`
	Zone       string `json:"zone,omitempty"`
	InstanceID string `json:"instance_id,omitempty"`
}

type AssetRatings struct {
	ACR *RatingScore `json:"acr,omitempty"`
	AES *RatingScore `json:"aes,omitempty"`
}

type RatingScore struct {
	Score float64 `json:"score"`
}
