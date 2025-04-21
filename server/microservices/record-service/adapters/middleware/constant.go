package middleware

const (
	CreateRecord = "/RecordService/CreateRecord"
	GetRecords   = "/RecordService/GetRecords"
	GetRecord    = "/RecordService/GetRecord"
	ScanUpload   = "/RecordService/ScanUpload"
	SearchRecord = "/RecordService/SearchRecord"
)

type UserType struct {
	UserID string
	Type   string
}
