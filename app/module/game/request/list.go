package request

type List struct {
	Name string `json:"name" example:"test"`
	From int64  `json:"from" example:"0" validate:"required,min=0"`
	Size int64  `json:"size" example:"20" validate:"required,min=0"`
}
