package selenium

type Reply struct {
	SessionId string                 `json:"sessionId,omitempty"`
	Status    int                    `json:"status,omitempty"`
	Value     map[string]interface{} `json:"value,omitempty"`
}
