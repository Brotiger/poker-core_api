package response

type Error400 struct {
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}
