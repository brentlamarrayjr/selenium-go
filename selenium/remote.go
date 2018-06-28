package selenium

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"fmt"
)

type RemoteWebDriver struct {
	Session      Session
	Capabilities capabilities
	URL          string
}

//NewRemote returns a pointer to an implementation of the W3C WebDriver client protocol
func NewRemote(url string, caps capabilities) *RemoteWebDriver {

	return &RemoteWebDriver{Capabilities: caps, URL: url}

}

//NewSession creates a single instantiation of a particular user agent and returns the session ID.
func (wd *RemoteWebDriver) NewSession() (Session, error) {

	//jsonData, err := json.Marshal(map[string]interface{}{"capabilities": map[string]interface{}{"alwaysMatch": capabilities}})
	jsonData, err := json.Marshal(map[string]interface{}{"desiredCapabilities": wd.Capabilities})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(wd.URL+"/session", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	session, err := ParseSession(body)
	if err != nil {
		return nil, err
	}

	wd.Session = session

	return session, nil

}

func (wd *RemoteWebDriver) GetStatus() (status, error) {
	resp, err := http.Get(fmt.Sprintf("%s/status", wd.URL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	status := &Status{Raw: body}

	w3cStatus, err := status.GetW3CStatus()
	if err == nil {
		return w3cStatus, nil
	}

	ncStatus, err := status.GetNCStatus()
	if err != nil {
		return nil, err
	}

	return ncStatus, nil

}

func (wd *RemoteWebDriver) DeleteSession() error {

	// Create client
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/session/%s", wd.URL, wd.Session.GetID()), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorCheck(resp)

}

func errorCheck(response *http.Response) error {

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	remoteErr := struct {
		*Reply
		Error *Error `json:"value,omitempty"`
	}{}

	err = json.Unmarshal(body, &remoteErr)
	if err != nil {

		return err
	}

	if remoteErr.Error != nil && remoteErr.Error.Message != "" {
		return errors.New(remoteErr.Error.Message)
	}

	if response.StatusCode != 200 {
		return errors.New("Invalid HTTP request status: " + response.Status)
	}

	return nil

}

func (wd *RemoteWebDriver) Navigate(url string) error {

	jsonData, _ := json.Marshal(map[string]interface{}{"url": url})

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/url", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorCheck(resp)

}

func (wd *RemoteWebDriver) GetCurrentURL() (string, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/url", wd.URL, wd.Session.GetID()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	url := ""
	err = json.Unmarshal(body, &url)
	if err != nil {
		return "", err
	}

	return url, nil

}

func (wd *RemoteWebDriver) Back() error {

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/back", wd.URL, wd.Session.GetID()), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorCheck(resp)

}

func (wd *RemoteWebDriver) Forward() error {

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/forward", wd.URL, wd.Session.GetID()), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorCheck(resp)

}

func (wd *RemoteWebDriver) Refresh() error {

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/refresh", wd.URL, wd.Session.GetID()), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorCheck(resp)

}

func (wd *RemoteWebDriver) GetTitle() (string, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/title", wd.URL, wd.Session.GetID()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	title := ""
	err = json.Unmarshal(body, &title)
	if err != nil {
		return "", err
	}

	return title, nil

}

func (wd *RemoteWebDriver) GetWindowHandle() (string, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/window", wd.URL, wd.Session.GetID()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	window := ""
	err = json.Unmarshal(body, &window)
	if err != nil {
		return "", err
	}

	return window, nil

}

func (wd *RemoteWebDriver) CloseWindow() error {

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/session/%s/window", wd.URL, wd.Session.GetID()), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorCheck(resp)

}

func (wd *RemoteWebDriver) SwitchToWindow(window string) error {

	jsonData, _ := json.Marshal(map[string]interface{}{"handle": window})
	resp, err := http.Post(fmt.Sprintf("%s/session/%s/window", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorCheck(resp)

}

func (wd *RemoteWebDriver) GetWindowHandles() ([]string, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/window/handles", wd.URL, wd.Session.GetID()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	window := make([]string, 0)
	err = json.Unmarshal(body, window)
	if err != nil {
		return nil, err
	}

	return window, nil

}

func (wd *RemoteWebDriver) SwitchToFrame(id int) error {

	jsonData, _ := json.Marshal(map[string]interface{}{"id": id})
	resp, err := http.Post(fmt.Sprintf("%s/session/%s/frame", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (wd *RemoteWebDriver) SwitchToParentFrame() error {

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/frame/parent", wd.URL, wd.Session.GetID()), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (wd *RemoteWebDriver) GetWindowRect() (map[string]interface{}, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/window/rect", wd.URL, wd.Session.GetID()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	windowRect := make(map[string]interface{})
	err = json.Unmarshal(body, windowRect)
	if err != nil {
		return nil, err
	}

	return windowRect, nil

}

func (wd *RemoteWebDriver) SetWindowRect(width int, height int, x int, y int) error {

	jsonData, _ := json.Marshal(map[string]interface{}{
		"width":  width,
		"height": height,
		"x":      x,
		"y":      y,
	})

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/window/rect", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (wd *RemoteWebDriver) MaximizeWindow() error {

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/window/maximize", wd.URL, wd.Session.GetID()), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (wd *RemoteWebDriver) MinimizeWindow() error {

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/window/minimize", wd.URL, wd.Session.GetID()), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (wd *RemoteWebDriver) FullscreenWindow() error {

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/window/fullscreen", wd.URL, wd.Session.GetID()), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (wd *RemoteWebDriver) FindElement(by By, selector string) (map[string]interface{}, error) {

	jsonData, _ := json.Marshal(map[string]interface{}{
		"using": by,
		"value": selector,
	})

	fmt.Println(string(jsonData))

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/element", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	err = errorCheck(resp)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Raw element: " + string(body))

	result := make(map[string]interface{})

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (wd *RemoteWebDriver) FindElements(by By, selector string) ([]map[string]interface{}, error) {

	jsonData, _ := json.Marshal(map[string]interface{}{
		"using": by,
		"value": selector,
	})

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/elements", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)

	err = json.Unmarshal(body, results)
	if err != nil {
		return nil, err
	}

	return results, nil

}

func (wd *RemoteWebDriver) FindElementFromElement(by By, selector string, id string) (map[string]interface{}, error) {

	jsonData, _ := json.Marshal(map[string]interface{}{
		"using": by,
		"value": selector,
	})

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/element/%s/element", wd.URL, wd.Session.GetID(), id), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (wd *RemoteWebDriver) FindElementsFromElement(by By, selector string, id string) ([]map[string]interface{}, error) {

	jsonData, _ := json.Marshal(map[string]interface{}{
		"using": by,
		"value": selector,
	})

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/elements/%s/elements", wd.URL, wd.Session.GetID(), id), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)

	err = json.Unmarshal(body, results)
	if err != nil {
		return nil, err
	}

	return results, nil

}

func (wd *RemoteWebDriver) GetActiveElement() (map[string]interface{}, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/active", wd.URL, wd.Session.GetID()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (wd *RemoteWebDriver) IsElementSelected(by By, selector string, id string) ([]map[string]interface{}, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/selected", wd.URL, wd.Session.GetID(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)

	err = json.Unmarshal(body, results)
	if err != nil {
		return nil, err
	}

	return results, nil

}
