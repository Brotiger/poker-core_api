package request

type Join struct {
	Password *string `json:"password,omitempty" example:"123456"`
}
