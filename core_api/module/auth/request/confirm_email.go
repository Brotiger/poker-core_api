package request

type ConfirmedEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
