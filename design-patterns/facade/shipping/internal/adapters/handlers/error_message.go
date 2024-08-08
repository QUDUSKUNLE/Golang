package handlers

type ErrorMessage string

const (
	USER_ALREADY_EXIST ErrorMessage = "user`s already exist"
	USER_ALREADY_REGISTERED ErrorMessage = "User`s already regsitered"
	INCORRECT_PASSWORDS ErrorMessage = "incorrect passwords"
)
