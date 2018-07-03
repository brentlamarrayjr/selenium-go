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

func (wd *RemoteWebDriver) GetTimeouts() (*Timeouts, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/timeouts", wd.URL, wd.Session.GetID()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	err = errorCheck(resp)
	if err != nil {
		return nil, err
	}

	reply := new(Reply)

	err = json.Unmarshal(body, reply)
	if err != nil {
		return nil, err
	}

	script, found := reply.Value["script"].(float64)
	if !found {
		return nil, errors.New("Missing Script timeout")
	}

	pageLoad, found := reply.Value["pageLoad"].(float64)
	if !found {
		return nil, errors.New("Missing PageLoad timeout")
	}

	implicit, found := reply.Value["implicit"].(float64)
	if !found {
		return nil, errors.New("Missing Implicit timeout")
	}

	return &Timeouts{Script: int(script), PageLoad: int(pageLoad), Implicit: int(implicit)}, nil

}

func (wd *RemoteWebDriver) SetTimeouts(timeouts *Timeouts) error {

	jsonData, err := json.Marshal(timeouts)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/timeouts", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return errors.New(message)
	} else if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil

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

	if len(body) > 0 {

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

func (wd *RemoteWebDriver) FindElement(by By, selector string) (WebElement, error) {

	jsonData, err := json.Marshal(map[string]interface{}{
		"using": by,
		"value": selector,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/element", wd.URL, wd.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return nil, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return nil, errors.New(message)
	} else if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	for _, value := range reply.Value {
		return &webElement{id: value.(string), driver: wd}, nil
	}

	return nil, errors.New("No element found")

}

func (wd *RemoteWebDriver) FindElements(by By, selector string) ([]WebElement, error) {

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

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return nil, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return nil, errors.New(message)
	} else if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	elements := make([]WebElement, 0)

	for _, value := range reply.Value {
		for _, value2 := range value.(map[string]interface{}) {
			elements = append(elements, &webElement{id: value2.(string), driver: wd})
		}
	}

	return elements, nil

}

func (wd *RemoteWebDriver) FindElementFromElement(by By, selector string, id string) (WebElement, error) {

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

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return nil, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return nil, errors.New(message)
	} else if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	for _, value := range reply.Value {
		return &webElement{id: value.(string), driver: wd}, nil
	}

	return nil, errors.New("No element found")

}

func (wd *RemoteWebDriver) FindElementsFromElement(by By, selector string, id string) ([]WebElement, error) {

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

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return nil, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return nil, errors.New(message)
	} else if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	elements := make([]WebElement, 0)

	for _, value := range reply.Value {
		for _, value2 := range value.(map[string]interface{}) {
			elements = append(elements, &webElement{id: value2.(string), driver: wd})
		}
	}

	return elements, nil
}

func (wd *RemoteWebDriver) GetActiveElement() (WebElement, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/active", wd.URL, wd.Session.GetID()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return nil, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return nil, errors.New(message)
	} else if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	for _, value := range reply.Value {
		return &webElement{id: value.(string), driver: wd}, nil
	}

	return nil, errors.New("No element found")

}

func (wd *RemoteWebDriver) IsElementSelected(element WebElement) (bool, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/selected", wd.URL, wd.Session.GetID(), element.GetWebDriverID()))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return false, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return false, errors.New(message)
	} else if resp.StatusCode != 200 {
		return false, errors.New(resp.Status)
	}

	selected, found := reply.Value["selected"].(bool)
	if found {
		return selected, nil
	}

	return false, errors.New("Attribute 'selected' not found")

}

func (wd *RemoteWebDriver) IsElementEnabled(element WebElement) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/enabled", wd.URL, wd.Session.GetID(), element.GetWebDriverID()))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return false, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return false, errors.New(message)
	} else if resp.StatusCode != 200 {
		return false, errors.New(resp.Status)
	}

	selected, found := reply.Value["selected"].(bool)
	if found {
		return selected, nil
	}

	return false, errors.New("Attribute 'selected' not found")
}
func (wd *RemoteWebDriver) GetElementAttribute(element WebElement, name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", wd.URL, wd.Session.GetID(), element.GetWebDriverID(), name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return "", err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return "", errors.New(message)
	} else if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	result, found := reply.Value["result"].(string)
	if found {
		return result, nil
	}

	return "", errors.New("Attribute 'result' not found")
}
func (wd *RemoteWebDriver) GetElementProperty(element WebElement, name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", wd.URL, wd.Session.GetID(), element.GetWebDriverID(), name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return "", err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return "", errors.New(message)
	} else if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	result, found := reply.Value["result"].(string)
	if found {
		return result, nil
	}

	return "", errors.New("Attribute 'result' not found")
}
func (wd *RemoteWebDriver) GetElementCSS(element WebElement, name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", wd.URL, wd.Session.GetID(), element.GetWebDriverID(), name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return "", err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return "", errors.New(message)
	} else if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	value, found := reply.Value["computed value"].(string)
	if found {
		return value, nil
	}

	return "", errors.New("Attribute 'computed value' not found")
}
func (wd *RemoteWebDriver) GetElementText(element WebElement) (string, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/text", wd.URL, wd.Session.GetID(), element.GetWebDriverID()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return "", err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return "", errors.New(message)
	} else if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	text, found := reply.Value["rendered text"].(string)
	if found {
		return text, nil
	}

	return "", errors.New("Attribute 'rendered text' not found")

}
func (wd *RemoteWebDriver) GetElementTagName(element WebElement) (string, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/name", wd.URL, wd.Session.GetID(), element.GetWebDriverID()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return "", err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return "", errors.New(message)
	} else if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	result, found := reply.Value["qualified name"].(string)
	if found {
		return result, nil
	}

	return "", errors.New("Attribute 'qualified name' not found")

}
func (wd *RemoteWebDriver) GetElementRect(element WebElement) (*Rect, error) {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/rect", wd.URL, wd.Session.GetID(), element.GetWebDriverID()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return nil, err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return nil, errors.New(message)
	} else if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	x, found := reply.Value["x"].(int)
	if !found {
		return nil, errors.New("'x' not found")
	}

	y, found := reply.Value["y"].(int)
	if !found {
		return nil, errors.New("'y' not found")
	}

	width, found := reply.Value["width"].(int)
	if !found {
		return nil, errors.New("'width' not found")
	}

	height, found := reply.Value["height"].(int)
	if !found {
		return nil, errors.New("'height' not found")
	}

	return &Rect{X: x, Y: y, Width: width, Height: height}, nil

}
func (wd *RemoteWebDriver) ElementClick(element WebElement) error {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/click", wd.URL, wd.Session.GetID(), element.GetWebDriverID()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return errors.New(message)
	} else if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil

}
func (wd *RemoteWebDriver) ElementClear(element WebElement) error {

	resp, err := http.Get(fmt.Sprintf("%s/session/%s/element/%s/clear", wd.URL, wd.Session.GetID(), element.GetWebDriverID()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return errors.New(message)
	} else if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil

}
func (wd *RemoteWebDriver) ElementSendKeys(element WebElement, keys string) error {

	jsonData, _ := json.Marshal(map[string]interface{}{
		"text": keys,
	})

	resp, err := http.Post(fmt.Sprintf("%s/session/%s/element/%s/value", wd.URL, wd.Session.GetID(), element.GetWebDriverID), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reply := new(Reply)
	err = json.Unmarshal(body, reply)
	if err != nil {
		return err
	}

	message, found := reply.Value["message"].(string)
	if found {
		return errors.New(message)
	} else if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil

}
