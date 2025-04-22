package middleware

const (
	CreateRecord = "/RecordService/CreateRecord"
	GetRecords   = "/RecordService/GetRecords"
	GetRecord    = "/RecordService/GetRecord"
	ScanUpload   = "/RecordService/ScanUpload"
	SearchRecord = "/RecordService/SearchRecord"
	SearchByNin  = "/RecordService/SearchByNin"
)

type UserType struct {
	UserID string
	Type   string
}
