package handlers

type ResponseMessage string

const (
	USER_REGISTERED_SUCCESSFULLY ResponseMessage = "User registered successfully."
	RESET_EMAIL_SENT_SUCCESSFULLY ResponseMessage = "Reset email sent successfully."
	// Shipping
	PRODUCT_SCHEDULED_FOR_SHIPPING ResponseMessage = "Product is scheduled for shipping."
	PRODUCT_IS_DELIVERED ResponseMessage = "Product is delivered."
	UPDATE_PARCEL_SUCCESSFULLY ResponseMessage = "Update parcel successfully."
	ADDRESSES_SUBMITTED_SUCCESSFULLY ResponseMessage = "User addresses submitted successfully."
)
