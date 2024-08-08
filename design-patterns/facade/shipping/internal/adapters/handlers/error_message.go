package handlers

type ErrorMessage string

const (
	USER_ALREADY_EXIST ErrorMessage = "user`s already exist"
	USER_ALREADY_REGISTERED ErrorMessage = "User`s already regsitered"
	INCORRECT_PASSWORDS ErrorMessage = "incorrect passwords"

	UNAUTHORIZED_TO_PERFORM_OPERATION ErrorMessage = "Unauthorized to perform this operation"
	RECORD_NOT_FOUND ErrorMessage = "record not found"
)
