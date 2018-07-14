package errors

import (
	"fmt"
	"time"
)

type seleniumError struct {
	when time.Time
	what string
}

func SeleniumError(message string) *seleniumError {
	return &seleniumError{time.Now(), message}
}

func (e *seleniumError) Error() string {
	return fmt.Sprintf("%v: %v", e.when, e.what)
}
