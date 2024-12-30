package request

type Restore struct {
	Email string `json:"email" example:"example@example.com" validate:"required,email"`
}
