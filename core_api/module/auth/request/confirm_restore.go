package request

type ConfirmedRestore struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}
