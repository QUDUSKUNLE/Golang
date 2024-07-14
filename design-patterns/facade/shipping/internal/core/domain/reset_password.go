package domain

type ResetPasswordDto struct {
	Email string `json:"Email" binding:"required" validate:"required,email"`
}
