package selenium

type Timeouts struct {
	Script   int `json:"script,omitempty"`
	PageLoad int `json:"pageLoad,omitempty"`
	Implicit int `json:"implicit,omitempty"`
}
