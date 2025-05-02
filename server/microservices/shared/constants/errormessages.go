package constants

// General error messages used across multiple services
const (
	ErrUnauthorized       = "Unauthorized to perform operation."
	ErrInvalidCredentials = "Incorrect login credentials."
)

// UserService error messages
const (
	ErrUserNotFound      = "User not found."
	ErrUserAlreadyExists = "User already exists."
	ErrUserNotRegistered = "User not registered."
	ErrInvalidUserType   = "Invalid user type."
	ErrInvalidNin        = "Invalid NIN provided."
	ErrInvalidEmail      = "Invalid email provided."
	ErrInvalidPassword   = "Invalid password provided."
)

// OrganizationService error messages
const (
	ErrInvalidOrganization = "Invalid organization provided."
)

// DiagnosticService error messages
const (
	ErrDiagnosticCentreNotFound = "Diagnostic centre not found."
)

// RecordService error messages
const (
	ErrRecordNotFound      = "Record not found."
	ErrRecordAlreadyExists = "Record already exists."
)
