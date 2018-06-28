package selenium

type Error struct {
	Error      string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
	Stacktrace string `json:"stacktrace,omitempty"`
}
