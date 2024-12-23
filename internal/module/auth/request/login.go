package request

type Login struct {
	Username string `json:"username" example:"username" validate:"required"`
	Password string `json:"password" example:"password" validate:"required"`
}
