package domain

type (
	RecordDto struct {
		UserID         string `json:"user_id"`
		OrganizationID string `json:"organization_id"`
		Record         string `json:"record"`
	}
	UploadDto struct {
		UserID         string `json:"user_id"`
		OrganizationID string `json:"organization_id"`
		ScanTitle      string `json:"scan_title"`
		FileName       string `json:"file_name"`
	}
)