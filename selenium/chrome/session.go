package chrome

type session struct {
	ID           string
	Capabilities map[string]interface{}
}

func (session *session) GetID() string {
	return session.ID
}

func (session *session) GetCapabilities() map[string]interface{} {
	return session.Capabilities
}
