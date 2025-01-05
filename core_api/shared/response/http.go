package response

type BadRequest struct {
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

type Unauthorized struct {
	Message string `json:"message"`
}

type NotFound struct {
	Message string `json:"message"`
}

type OK struct {
	Message string `json:"message"`
}

type Forbidden struct {
	Message string `json:"message"`
}
