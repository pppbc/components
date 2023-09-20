package errors

import "errors"

var (
	// ErrThird.
	ErrThird = errors.New("三方服务异常")
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
)
