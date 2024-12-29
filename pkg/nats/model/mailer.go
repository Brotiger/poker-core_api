package model

type Mailer struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Type  string `json:"type"`
}
