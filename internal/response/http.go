package response

type Error400 struct {
	Message string `json:"message" default:"Bad request"`
	Errors  any    `json:"errors"`
}
