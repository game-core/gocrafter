package errors

import (
	"fmt"
)

// NewError エラー
func NewError(method string, err error) error {
	return fmt.Errorf("failed to %s: %s", method, err)
}

// NewTestError テストエラー
func NewTestError(method string, err error) error {
	return fmt.Errorf("test error")
}
