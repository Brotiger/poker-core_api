package response

type Error400 struct {
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

type Error401 struct {
	Message string `json:"message"`
}

type Error404 struct {
	Message string `json:"message"`
}
