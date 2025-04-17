package middleware

const (
	CreateOrganization = "/OrganizationService/CreateOrganization"
	GetOrganization    = "/OrganizationService/GetOrganization"
)

type UserType struct {
	UserID string
	Type   string
}
