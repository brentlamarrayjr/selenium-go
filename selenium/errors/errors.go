package errors

func WebDriverInfoNotImplemented() *seleniumError {
	return SeleniumError("WebDriverInfo not implemented")
}

func WebElementInfoNotImplemented() *seleniumError {
	return SeleniumError("WebElementInfo not implemented")
}

func WebElementUpdaterNotImplemented() *seleniumError {
	return SeleniumError("WebElementUpdater not implemented")
}
