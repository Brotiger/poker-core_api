package request

type Register struct {
	Email    string `json:"email" example:"example@example.com" validate:"required,email"`
	Username string `json:"username" example:"example" validate:"required"`
	Password string `json:"password" example:"example" validate:"required"`
}
