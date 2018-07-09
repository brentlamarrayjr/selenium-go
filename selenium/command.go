package selenium

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	DELETE Method = "DELETE"
)

func ExecuteWDCommand(method Method, endpoint string, data interface{}) (*Reply, error) {

	var resp *http.Response
	var err error

	if method == POST {

		if data != nil {
			jsonData, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			resp, err = http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
		} else {
			resp, err = http.Post(endpoint, "", nil)
		}

	} else if method == GET {
		resp, err = http.Get(endpoint)
	} else if method == DELETE {
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", endpoint, nil)
		if err != nil {
			return nil, err
		}

		resp, err = client.Do(req)

	} else {
		return nil, errors.New("unsupported http method: " + string(method))
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	replyData := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &replyData)
	if err != nil {
		return nil, err
	}

	return &Reply{StatusCode: resp.StatusCode, Data: replyData}, nil

}
