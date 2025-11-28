package errors

import (
	"errors"
	"fmt"

	pkgerr "github.com/pkg/errors"
)

var errMap = map[int]Error{}

type Error struct {
	code    int
	message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.code, e.message)
}

func NewError(code int, message string) Error {
	if _, ok := errMap[code]; ok {
		panic(fmt.Sprintf("错误码：%d 已定义", code))
	}
	e := Error{
		code:    code,
		message: message,
	}
	errMap[code] = e
	return e
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Message() string {
	return e.message
}

func Is(err error, target error) bool {
	return pkgerr.Is(err, target)
}

func As(err error, target error) bool {
	return pkgerr.As(err, target)
}

func Wrap(err error, msg string) error {
	return pkgerr.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...any) error {
	return pkgerr.Wrapf(err, format, args...)
}

func Unwrap(err error) error {
	return pkgerr.Unwrap(err)
}

func WithMessage(err error, message string) error {
	return pkgerr.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...any) error {
	return pkgerr.WithMessagef(err, format, args...)
}

func WithStack(err error) error {
	return pkgerr.WithStack(err)
}

func Cause(err error) error {
	return pkgerr.Cause(err)
}

func Errorf(format string, args ...any) error {
	return pkgerr.Errorf(format, args...)
}

func Join(err ...error) error {
	return errors.Join(err...)
}

func New(text string) error {
	return errors.New(text)
}
