package request

type Login struct {
	Email    string `json:"email" example:"example@example.com" validate:"required,email"`
	Password string `json:"password" example:"password" validate:"required"`
}
