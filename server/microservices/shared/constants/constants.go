package constants

type UserType struct {
	UserID string
	Type   string
}

const (
	// Users constants
	UpdateUser = "/UserService/UpdateUser"
	ReadUsers  = "/UserService/ReadUsers"

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

	// Diagnostic case
	CreateDiagnostic         = "/DiagnosticService/CreateDiagnostic"
	GetDiagnostic            = "/DiagnosticService/GetDiagnostic"
	ListDiagnostics          = "/DiagnosticService/ListDiagnostics"
	DeleteDiagnostic         = "/DiagnosticService/DeleteDiagnostic"
	UpdateDiagnostic         = "/DiagnosticService/UpdateDiagnostic"
	SearchNearestDiagnostics = "/DiagnosticService/SearchNearestDiagnostics"

	// Schedule case
	CreateSchedule        = "/ScheduleService/CreateScheduleSession"
	GetScheduleSession    = "/ScheduleService/GetScheduleSession"
	ListScheduleSessions  = "/ScheduleService/ListScheduleSessions"
	DeleteScheduleSession = "/ScheduleService/DeleteScheduleSession"
	UpdateScheduleSession = "/ScheduleService/UpdateScheduleSession"

	DefaultLimit  = 50
	DefaultOffset = 0
)
