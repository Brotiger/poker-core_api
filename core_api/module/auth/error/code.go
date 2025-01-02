package error

import "errors"

var ErrCodeNotFound error
var ErrCompareCode error

func init() {
	ErrCodeNotFound = errors.New("code not found")
	ErrCompareCode = errors.New("failed compare code")
}
