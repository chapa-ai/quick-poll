package errors

import (
	"errors"
	"fmt"
)

var (
	ErrPollNotFound  = errors.New("poll not found")
	ErrInvalidOption = errors.New("invalid option")
)

func New(msg string) error {
	return errors.New(msg)
}

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

func NewF(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}
