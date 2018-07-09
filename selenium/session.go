package selenium

type Session struct {
	ID           string
	Capabilities map[string]interface{}
}

func (session *Session) GetID() string {
	return session.ID
}

func (session *Session) GetCapabilities() map[string]interface{} {
	return session.Capabilities
}
