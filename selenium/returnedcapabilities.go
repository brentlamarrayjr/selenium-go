package selenium

//Capabilities - Selenium capabilities
type ReturnedCapabilities struct {
	BrowserName             string            `json:"browserName,omitempty"`
	BrowserVersion          string            `json:"browserVersion,omitempty"`
	PlatformName            string            `json:"platformName,omitempty"`
	AcceptInsecureCerts     bool              `json:"acceptInsecureCerts,omitempty"`
	PageLoadStrategy        *PageLoadStrategy `json:"pageLoadStrategy,omitempty"`
	Proxy                   *Proxy            `json:"proxy,omitempty"`
	WindowRect              bool              `json:"setWindowRect,omitempty"`
	Timeouts                *Timeouts         `json:"timeouts,omitempty"`
	UnhandledPromptBehavior string            `json:"unhandledPromptBehavior,omitempty"`
	Rotatable               bool              `json:"rotatable,omitempty"`

	//Chrome
	Chrome                   map[string]interface{} `json:"chrome,omitempty"`
	CSSSelectorsEnabled      bool                   `json:"cssSelectorsEnabled,omitempty"`
	DatabaseEnabled          bool                   `json:"databaseEnabled,omitempty"`
	HandlesAlerts            bool                   `json:"handlesAlerts,omitempty"`
	HasTouchScreen           bool                   `json:"hasTouchScreen,omitempty"`
	JavascriptEnabled        bool                   `json:"javascriptEnabled,omitempty"`
	LocationContextEnabled   bool                   `json:"locationContextEnabled,omitempty"`
	MobileEmulationEnabled   bool                   `json:"mobileEmulationEnabled,omitempty"`
	NativeEvents             bool                   `json:"nativeEvents,omitempty"`
	NetworkConnectionEnabled bool                   `json:"networkConnectionEnabled,omitempty"`
	Platform                 string                 `json:"platform,omitempty"`
	Version                  string                 `json:"version,omitempty"`
	TakesHeapSnapshot        bool                   `json:"takesHeapSnapshot,omitempty"`
	TakesScreenshot          bool                   `json:"takesScreenshot,omitempty"`
	UnhandledAlertBehavior   string                 `json:"unhandledAlertBehavior,omitempty"`
	AcceptSSLCerts           bool                   `json:"acceptSslCerts,omitempty"`
	ApplicationCacheEnabled  bool                   `json:"applicationCacheEnabled,omitempty"`
	BrowserConnectionEnabled bool                   `json:"browserConnectionEnabled,omitempty"`

	//Firefox
	AccessibilityChecks              bool
	Headless                         bool
	ProcessID                        string
	Profile                          string
	UseNonSpecCompliantPointerOrigin bool
	WebDriverClick                   bool
}
