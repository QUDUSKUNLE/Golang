package constants

type UserType struct {
	UserID string
	Type   string
}

const (
	// Users constants
	UpdateNin = "/UserService/UpdateNin"
	ReadUsers = "/UserService/ReadUsers"
	// Records constants
	CreateRecord = "/RecordService/CreateRecord"
	GetRecords   = "/RecordService/GetRecords"
	GetRecord    = "/RecordService/GetRecord"
	ScanUpload   = "/RecordService/ScanUpload"
	SearchRecord = "/RecordService/SearchRecord"
	SearchByNin  = "/RecordService/SearchByNin"

	// Organization case
	CreateOrganization = "/OrganizationService/CreateOrganization"
	GetOrganization    = "/OrganizationService/GetOrganization"
)
