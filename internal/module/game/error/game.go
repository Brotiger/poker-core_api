package error

import "errors"

var ErrGameNotFound error

func init() {
	ErrGameNotFound = errors.New("game not found")
}
