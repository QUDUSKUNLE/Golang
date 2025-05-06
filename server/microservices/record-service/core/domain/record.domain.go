package domain

type (
	RecordDto struct {
		UserID         string `json:"user_id"`
		DiagnosticID string `json:"diagnostic_id"`
		Record         string `json:"record"`
		ScanTitle      string `json:"scan_title"`
	}
	UploadDto struct {
		UserID         string `json:"user_id"`
		ScanTitle      string `json:"scan_title"`
		FileName       string `json:"file_name"`
		DiagnosticID string `json:"diagnostic_id"`
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
