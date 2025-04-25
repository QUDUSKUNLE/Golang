package middleware

const (
	UpdateNin = "/UserService/UpdateNin"
	ReadUsers = "/UserService/ReadUsers"
)

type UserType struct {
	UserID string
	Type   string
}
