package chrome

import "../../selenium"

type Capabilities struct {
	selenium.Capabilities
	ChromeOptions *ChromeOptions `json:"chromeOptions,omitempty"`
}
