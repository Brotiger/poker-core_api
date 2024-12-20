package validatorError

import "fmt"

type Min struct{}

var DefaultError = "Неизвестная ошибка."

func (min *Min) GetError(param string) string {
	return fmt.Sprintf("Минимальная длина: %s.", param)
}
