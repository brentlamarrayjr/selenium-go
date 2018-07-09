package selenium

type Status struct {
	Ready   bool   `json:"ready,omitempty"`
	Message string `json:"message,omitempty"`
}

func (status *Status) IsReady() bool {
	return status.Ready
}

func (status *Status) GetMessage() string {
	return status.Message
}
