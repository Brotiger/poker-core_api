package request

type List struct {
	Name string `json:"name"`
	From int64  `json:"from"`
	Size int64  `json:"size"`
}
