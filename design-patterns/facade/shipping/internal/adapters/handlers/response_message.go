package handlers

type ResponseMessage string

const (
	USER_REGISTERED_SUCCESSFULLY ResponseMessage = "User registered successfully."
	RESET_EMAIL_SENT_SUCCESSFULLY ResponseMessage = "Reset email sent successfully."
)
