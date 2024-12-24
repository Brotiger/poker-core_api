package validatorError

import "fmt"

type Min struct{}
type Required struct{}

var DefaultError = "Неизвестная ошибка."

func (min *Min) GetError(param string) string {
	return fmt.Sprintf("Минимальная длина: %s.", param)
}

func (required *Required) GetError(param string) string {
	return "Обязательно для заполнения."
}
