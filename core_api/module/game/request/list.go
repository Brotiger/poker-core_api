package request

type List struct {
	Name *string `json:"name,omitempty" example:"test"`
	From int64   `json:"from" example:"0" validate:"min=0"`
	Size int64   `json:"size" example:"20" validate:"required,min=0"`
}
