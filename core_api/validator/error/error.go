package validatorError

import "fmt"

type Min struct{}
type Required struct{}
type Email struct{}

var DefaultError = "Неизвестная ошибка."

func (min *Min) GetError(param string) string {
	return fmt.Sprintf("Минимальная длина: %s.", param)
}

func (required *Required) GetError(param string) string {
	return "Обязательно для заполнения."
}

func (email *Email) GetError(param string) string {
	return "Невалидный формат."
}
