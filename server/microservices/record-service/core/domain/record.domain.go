package domain

type (
	RecordDto struct {
		UserID         string `json:"user_id"`
		OrganizationID string `json:"organization_id"`
		Record         string `json:"record"`
	}
)
