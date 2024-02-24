package errors

import (
	"fmt"
)

// NewError エラー
func NewError(message string) error {
	return fmt.Errorf(message)
}

// NewMethodError エラー
func NewMethodError(method string, err error) error {
	return fmt.Errorf("failed to %s: %s", method, err)
}

// NewTestError テストエラー
func NewTestError() error {
	return fmt.Errorf("test")
}
