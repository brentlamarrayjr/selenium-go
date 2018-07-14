package selenium

type WebDriverInfo interface {
	GetURL() string
	GetDesiredCapabilities() Capabilities
	GetSession() SessionInfo
}
