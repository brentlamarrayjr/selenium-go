package errors

import "time"

type seleniumError struct {
	when time.Time
	what string
}

func SeleniumError(message string) *seleniumError {
	return &seleniumError{time.Now(), message}
}
