package middleware

const (
	CreateRecord = "/RecordService/CreateRecord"
	GetRecords   = "/RecordService/GetRecords"
	GetRecord    = "/RecordService/GetRecord"
	ScanUpload   = "/RecordService/ScanUpload"
)

type UserType struct {
	UserID string
	Type   string
}
