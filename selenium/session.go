package selenium

import (
	"encoding/json"
	"errors"
)

type session struct {
	id           string
	capabilities *ReturnedCapabilities
}

type Session interface {
	GetID() string
	GetReturnedCapabilities() *ReturnedCapabilities
}

func ParseSession(data []byte) (Session, error) {

	if data == nil {
		return nil, errors.New("no session data")
	}

	sesh := struct {
		*Reply
		Session *struct {
			ID           string                `json:"sessionId,omitempty"`
			Capabilities *ReturnedCapabilities `json:"capabilities,omitempty"`
		} `json:"value,omitempty"`
		//Session *session `json:"value,omitempty"`
	}{}

	err := json.Unmarshal(data, &sesh)
	if err != nil {
		return nil, err
	}

	if sesh.Session.ID != "" {
		return &session{id: sesh.Session.ID, capabilities: sesh.Session.Capabilities}, nil
	}

	ncsesh := struct {
		*Reply
		Capabilities *ReturnedCapabilities `json:"value,omitempty"`
	}{}

	err = json.Unmarshal(data, &ncsesh)
	if err != nil {
		return nil, err
	}

	if ncsesh.SessionId == "" {
		return nil, errors.New("missing session ID")
	}

	return &session{id: ncsesh.SessionId, capabilities: ncsesh.Capabilities}, nil
}

func (session *session) GetID() string {
	return session.id
}

func (session *session) GetReturnedCapabilities() *ReturnedCapabilities {
	return session.capabilities
}
