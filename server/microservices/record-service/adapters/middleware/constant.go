package middleware

const (
	CreateRecord = "/RecordService/CreateRecord"
	GetRecords   = "/RecordService/GetRecords"
)

type UserType struct {
	UserID string
	Type   string
}
