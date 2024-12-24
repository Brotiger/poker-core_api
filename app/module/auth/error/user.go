package error

import "errors"

var ErrUserNotFound error
var ErrCompareHashAndPassword error

func init() {
	ErrUserNotFound = errors.New("user not found")
	ErrCompareHashAndPassword = errors.New("failed compare hash and password")
}
