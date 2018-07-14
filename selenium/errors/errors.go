package errors

import (
	"fmt"
)

func WebDriverInfoNotImplemented() *seleniumError {
	return SeleniumError("WebDriverInfo not implemented")
}

func WebElementInfoNotImplemented() *seleniumError {
	return SeleniumError("WebElementInfo not implemented")
}

func WebElementUpdaterNotImplemented() *seleniumError {
	return SeleniumError("WebElementUpdater not implemented")
}

func (e *seleniumError) Error() string {
	return fmt.Sprintf("%v: %v", e.when, e.what)
}
