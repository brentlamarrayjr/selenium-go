package selenium

//Capabilities - Selenium capabilities
type capabilities struct {
	BrowserName             string            `json:"browserName,omitempty"`
	BrowserVersion          string            `json:"browserVersion,omitempty"`
	PlatformName            string            `json:"platformName,omitempty"`
	AcceptInsecureCerts     bool              `json:"acceptInsecureCerts,omitempty"`
	PageLoadStrategy        *PageLoadStrategy `json:"pageLoadStrategy,omitempty"`
	Proxy                   *Proxy            `json:"proxy,omitempty"`
	WindowRect              bool              `json:"setWindowRect,omitempty"`
	Timeouts                *Timeouts         `json:"timeouts,omitempty"`
	UnhandledPromptBehavior string            `json:"unhandledPromptBehavior,omitempty"`
}

//Capabilities provides an inteface to common W3C driver capabilities
type Capabilities interface {
	SetBrowserName(name string)
	SetBrowserVersion(version string)
	SetPlatformName(name string)
	SetAcceptInsecureCerts(accept bool)
	SetPageLoadStrategy(pls *PageLoadStrategy)
	SetProxy(proxy *Proxy)
	SetWindowRect(wr bool)
	SetTimeouts(timeouts *Timeouts)
	SetUnhandledPromptBehavior(uhp string)
}

//NewCapabilities returns an implementation of the Capabilities interface
func NewCapabilities() Capabilities {
	return &capabilities{}
}

//SetBrowserName sets the lowercase name of the user agent
func (caps *capabilities) SetBrowserName(name string) {
	caps.BrowserName = name
}

//SetBrowserVersion sets the user agent version
func (caps *capabilities) SetBrowserVersion(version string) {
	caps.BrowserVersion = version
}

//SetPlatformName sets the operating system of the endpoint node
func (caps *capabilities) SetPlatformName(name string) {
	caps.PlatformName = name
}

//SetAcceptInsecureCerts  determines if the session will implicitly trust untrusted or self-signed TLS certificates on navigation.
func (caps *capabilities) SetAcceptInsecureCerts(accept bool) {
	caps.AcceptInsecureCerts = accept
}

//SetPageLoadStrategy sets the current session’s page load strategy.
func (caps *capabilities) SetPageLoadStrategy(pls *PageLoadStrategy) {
	caps.PageLoadStrategy = pls
}

//SetProxy sets the current session’s proxy configuration.
func (caps *capabilities) SetProxy(proxy *Proxy) {
	caps.Proxy = proxy
}

//SetWindowRect determines if the remote end supports all of the commands in Resizing and Positioning Windows.
func (caps *capabilities) SetWindowRect(wr bool) {
	caps.WindowRect = wr
}

//SetTimeouts sets the timeouts imposed on certain session operations.
func (caps *capabilities) SetTimeouts(timeouts *Timeouts) {
	caps.Timeouts = timeouts
}

//SetUnhandledPromptBehavior sets the current session’s user prompt handler.
func (caps *capabilities) SetUnhandledPromptBehavior(uhp string) {
	caps.UnhandledPromptBehavior = uhp
}
