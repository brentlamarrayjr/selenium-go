package selenium

//Capabilities - Selenium capabilities
type Capabilities struct {
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

type capabilities interface {
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

//SetBrowserName sets the lowercase name of the user agent
func (caps *Capabilities) SetBrowserName(name string) {
	caps.BrowserName = name
}

//SetBrowserVersion sets the user agent version
func (caps *Capabilities) SetBrowserVersion(version string) {
	caps.BrowserVersion = version
}

//SetPlatformName sets the operating system of the endpoint node
func (caps *Capabilities) SetPlatformName(name string) {
	caps.PlatformName = name
}

//SetAcceptInsecureCerts  determines if the session will implicitly trust untrusted or self-signed TLS certificates on navigation.
func (caps *Capabilities) SetAcceptInsecureCerts(accept bool) {
	caps.AcceptInsecureCerts = accept
}

//SetPageLoadStrategy sets the current session’s page load strategy.
func (caps *Capabilities) SetPageLoadStrategy(pls *PageLoadStrategy) {
	caps.PageLoadStrategy = pls
}

//SetProxy sets the current session’s proxy configuration.
func (caps Capabilities) SetProxy(proxy *Proxy) {
	caps.Proxy = proxy
}

//SetWindowRect determines if the remote end supports all of the commands in Resizing and Positioning Windows.
func (caps *Capabilities) SetWindowRect(wr bool) {
	caps.WindowRect = wr
}

//SetTimeouts sets the timeouts imposed on certain session operations.
func (caps *Capabilities) SetTimeouts(timeouts *Timeouts) {
	caps.Timeouts = timeouts
}

//SetUnhandledPromptBehavior sets the current session’s user prompt handler.
func (caps *Capabilities) SetUnhandledPromptBehavior(uhp string) {
	caps.UnhandledPromptBehavior = uhp
}
