package dto

type UserDTO struct {
	Email    				string `json:"email" binding:"required,email,lte=100" validate:"required"`
	Password 				string `json:"password" binding:"required,gte=6,lte=20" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,gte=6,lte=20" validate:"required"`
}
