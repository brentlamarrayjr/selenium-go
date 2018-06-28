package selenium

import (
	"encoding/json"
	"errors"
)

type status interface {
	IsReady() bool
}

type NCStatus struct {

	//chromedriver
	ID     string                 `json:"sessionId,omitempty"`
	Status int                    `json:"status,omitempty"`
	Value  map[string]interface{} `json:"value,omitempty"`
}

type W3CStatus struct {
	Ready   bool   `json:"ready,omitempty"`
	Message string `json:"message,omitempty"`
}

type Status struct {
	Raw []byte
}

func (status *Status) GetW3CStatus() (*W3CStatus, error) {

	if status.Raw == nil {
		return nil, errors.New("no status data")
	}

	w3cStatus := new(W3CStatus)

	err := json.Unmarshal(status.Raw, w3cStatus)
	if err != nil {
		return nil, err
	}

	if w3cStatus.Message == "" {
		return nil, errors.New("missing status message")
	}

	return w3cStatus, nil

}

func (status *Status) GetNCStatus() (*NCStatus, error) {

	if status.Raw == nil {
		return nil, errors.New("no status data")
	}

	ncStatus := new(NCStatus)

	err := json.Unmarshal(status.Raw, ncStatus)
	if err != nil {
		return nil, err
	}

	return ncStatus, nil

}

func (status *W3CStatus) IsReady() bool {

	return status.Ready
}

func (status *NCStatus) IsReady() bool {
	return status.Status == 0
}
