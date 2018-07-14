package selenium

type SessionInfo interface {
	GetID() string
	GetCapabilities() map[string]interface{}
}
