package error

import "errors"

var ErrGameNotFound error
var ErrComparePassword error

func init() {
	ErrGameNotFound = errors.New("game not found")
	ErrComparePassword = errors.New("failed to compare password")
}
