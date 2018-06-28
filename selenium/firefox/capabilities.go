package firefox

import "../../selenium"

type Capabilities struct {
	selenium.Capabilities
	FirefoxOptions *Options `json:"moz:firefoxOptions,omitempty"`
}
