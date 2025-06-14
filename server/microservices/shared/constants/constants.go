package constants

type UserType struct {
	UserID string
	Type   string
}

// Default values
const (
	DefaultLimit  = 50
	DefaultOffset = 0
	MaxLimit      = 100
	DistanceUnit  = "km"
)

// UserService constants
const (
	UpdateUser = "/UserService/UpdateUser"
	ReadUsers  = "/UserService/ReadUsers"
)

// RecordService constants
const (
	CreateRecord = "/RecordService/CreateRecord"
	GetRecords   = "/RecordService/GetRecords"
	GetRecord    = "/RecordService/GetRecord"
	ScanUpload   = "/RecordService/ScanUpload"
	SearchRecord = "/RecordService/SearchRecord"
	SearchByNin  = "/RecordService/SearchByNin"
)

// DiagnosticService constants
const (
	CreateDiagnostic         = "/DiagnosticService/CreateDiagnostic"
	GetDiagnostic            = "/DiagnosticService/GetDiagnostic"
	ListDiagnostics          = "/DiagnosticService/ListDiagnostics"
	DeleteDiagnostic         = "/DiagnosticService/DeleteDiagnostic"
	UpdateDiagnostic         = "/DiagnosticService/UpdateDiagnostic"
	SearchNearestDiagnostics = "/DiagnosticService/SearchNearestDiagnostics"
	ListDiagnosticSchedules  = "/DiagnosticService/ListDiagnosticSchedules"
	GetDiagnosticSchedule    = "/DiagnosticService/GetDiagnosticSchedule"
)

// ScheduleService constants
const (
	CreateSchedule = "/ScheduleService/CreateSchedule"
	GetSchedule    = "/ScheduleService/GetSchedule"
	ListSchedules  = "/ScheduleService/ListSchedules"
	CancelSchedule = "/ScheduleService/DeleteSchedule"
	UpdateSchedule = "/ScheduleService/UpdateSchedule"

	// Diagnostic Centre Schedule
	ListDiagnosticCentreSchedules = "/ScheduleService/ListDiagnosticCentreSchedules"
)
