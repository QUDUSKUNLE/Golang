package middleware

const (
	CreateRecord = "/RecordService/CreateRecord"
	GetRecords = "/RecordsService/GetRecords"
)

type UserType struct {
	UserID string
	Type   string
}
