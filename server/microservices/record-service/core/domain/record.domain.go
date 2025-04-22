package domain

type (
	RecordDto struct {
		UserID         string `json:"user_id"`
		OrganizationID string `json:"organization_id"`
		Record         string `json:"record"`
		ScanTitle      string `json:"scan_title"`
	}
	UploadDto struct {
		UserID         string `json:"user_id"`
		ScanTitle      string `json:"scan_title"`
		FileName       string `json:"file_name"`
		OrganizationID string `json:"organization_id"`
	}
	GetRecordDto struct {
		UserID    string `json:"user_id"`
		ScanTitle string `json:"scan_title"`
	}
	GetRecordByNinDto struct {
		Nin       string `json:"nin"`
		ScanTitle string `json:"scan_title"`
	}
)
