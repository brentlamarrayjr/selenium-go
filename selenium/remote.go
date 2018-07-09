package selenium

import (
	"errors"

	"fmt"
)

type RemoteWebDriver struct {
	Session      *Session
	Capabilities capabilities
	URL          string
}

//NewRemote returns a pointer to an implementation of the W3C WebDriver client protocol
func NewRemote(url string, caps capabilities) *RemoteWebDriver {

	return &RemoteWebDriver{Capabilities: caps, URL: url}

}

//NewSession creates a single instantiation of a particular user agent and returns the session ID.
func (wd *RemoteWebDriver) NewSession() (*Session, error) {

	//chrome map[string]interface{}{"desiredCapabilities": wd.Capabilities},

	reply, err := ExecuteWDCommand(
		POST,
		wd.URL+"/session",
		map[string]interface{}{
			"capabilities": map[string]interface{}{
				"alwaysMatch": wd.Capabilities,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	fmt.Println("Data: ")
	fmt.Println(reply.Data)

	id, err := reply.GetString("value.sessionId", true)
	if err != nil {
		return nil, err
	}

	capabilities, err := reply.GetMap("value.capabilities", true)
	if err != nil {
		return nil, err
	}

	fmt.Println("Capabilities: ")
	fmt.Println(capabilities)

	wd.Session = &Session{ID: id, Capabilities: capabilities}

	return wd.Session, nil

}

func (wd *RemoteWebDriver) GetStatus() (*Status, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/status", wd.URL),
		nil,
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	ready, err := reply.GetBool("value.ready", true)
	if err != nil {
		return nil, err
	}

	message, err := reply.GetString("value.message", true)
	if err != nil {
		return nil, err
	}

	return &Status{Ready: ready, Message: message}, nil

}

func (wd *RemoteWebDriver) GetTimeouts() (*Timeouts, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/timeouts", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	script, err := reply.GetFloat("value.script", true)
	if err != nil {
		return nil, err
	}

	pageLoad, err := reply.GetFloat("value.pageLoad", true)
	if err != nil {
		return nil, err
	}

	implicit, err := reply.GetFloat("value.implicit", true)
	if err != nil {
		return nil, err
	}

	return &Timeouts{Script: int(script), PageLoad: int(pageLoad), Implicit: int(implicit)}, nil

}

func (wd *RemoteWebDriver) SetTimeouts(timeouts *Timeouts) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/timeouts", wd.URL, wd.Session.GetID()),
		timeouts,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

//DeleteSession closes the current WebDriver session
func (wd *RemoteWebDriver) DeleteSession() error {

	reply, err := ExecuteWDCommand(
		DELETE,
		fmt.Sprintf("%s/session/%s", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) Navigate(url string) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/url", wd.URL, wd.Session.GetID()),
		map[string]interface{}{"url": url},
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) GetCurrentURL() (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/url", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	url, err := reply.GetString("value.url", true)
	if err != nil {
		return "", err
	}

	return url, nil

}

func (wd *RemoteWebDriver) Back() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/back", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) Forward() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/forward", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) Refresh() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/refresh", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) GetTitle() (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/title", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	title, err := reply.GetString("value.title", true)
	if err != nil {
		return "", err
	}

	return title, nil

}

func (wd *RemoteWebDriver) GetWindowHandle() (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/window", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	window, err := reply.GetString("value.window", true)
	if err != nil {
		return "", err
	}

	return window, nil

}

func (wd *RemoteWebDriver) CloseWindow() error {

	reply, err := ExecuteWDCommand(
		DELETE,
		fmt.Sprintf("%s/session/%s/window", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) SwitchToWindow(window string) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window", wd.URL, wd.Session.GetID()),
		map[string]interface{}{"handle": window},
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) GetWindowHandles() ([]string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/window/handles", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code")
	}

	handles, err := reply.GetStringSlice("value.handles", true)
	if err != nil {
		return nil, err
	}

	return handles, nil

}

func (wd *RemoteWebDriver) SwitchToFrame(id int) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/frame", wd.URL, wd.Session.GetID()),
		map[string]interface{}{"id": id},
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) SwitchToParentFrame() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/frame/parent", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) GetWindowRect() (*Rect, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/window/rect", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	x, err := reply.GetFloat("value.x", true)
	if err != nil {
		return nil, err
	}

	y, err := reply.GetFloat("value.y", true)
	if err != nil {
		return nil, err
	}

	width, err := reply.GetFloat("value.width", true)
	if err != nil {
		return nil, err
	}

	height, err := reply.GetFloat("value.height", true)
	if err != nil {
		return nil, err
	}

	return &Rect{X: int(x), Y: int(y), Width: int(width), Height: int(height)}, nil

}

func (wd *RemoteWebDriver) SetWindowRect(rect *Rect) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/rect", wd.URL, wd.Session.GetID()),
		rect,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) MaximizeWindow() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/maximize", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) MinimizeWindow() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/minimize", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) FullscreenWindow() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/fullscreen", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code")
	}

	return nil

}

func (wd *RemoteWebDriver) FindElement(by By, selector string) (WebElement, error) {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element", wd.URL, wd.Session.GetID()),
		map[string]interface{}{
			"using": by,
			"value": selector,
		},
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	elements, err := reply.GetStringMap("value", false)
	if err != nil {
		return nil, err
	}

	for id, value := range elements {
		return &webElement{id: id, value: value, driver: wd}, nil
	}

	return nil, errors.New("no element found")

}

func (wd *RemoteWebDriver) FindElements(by By, selector string) ([]WebElement, error) {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element", wd.URL, wd.Session.GetID()),
		map[string]interface{}{
			"using": by,
			"value": selector,
		},
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	elements := make([]WebElement, 0)

	foundElements, err := reply.GetStringMap("value", false)
	if err != nil {
		return nil, err
	}

	for id, value := range foundElements {
		elements = append(elements, &webElement{id: id, value: value, driver: wd})
	}

	return elements, nil

}

func (wd *RemoteWebDriver) FindElementFromElement(by By, selector string, element WebElement) (WebElement, error) {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/element", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		map[string]interface{}{
			"using": by,
			"value": selector,
		},
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	elements, err := reply.GetStringMap("value", false)
	if err != nil {
		return nil, err
	}

	for id, value := range elements {
		return &webElement{id: id, value: value, driver: wd}, nil
	}

	return nil, errors.New("no element found")

}

func (wd *RemoteWebDriver) FindElementsFromElement(by By, selector string, element WebElement) ([]WebElement, error) {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/element", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		map[string]interface{}{
			"using": by,
			"value": selector,
		},
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	elements := make([]WebElement, 0)

	foundElements, err := reply.GetStringMap("value", false)
	if err != nil {
		return nil, err
	}

	for id, value := range foundElements {
		elements = append(elements, &webElement{id: id, value: value, driver: wd})
	}

	return elements, nil
}

func (wd *RemoteWebDriver) GetActiveElement() (WebElement, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/active", wd.URL, wd.Session.GetID()),
		nil,
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	elements, err := reply.GetStringMap("value", false)
	if err != nil {
		return nil, err
	}

	for id, value := range elements {
		return &webElement{id: id, value: value, driver: wd}, nil
	}

	return nil, errors.New("no element found")

}

func (wd *RemoteWebDriver) IsElementSelected(element WebElement) (bool, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/selected", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		nil,
	)

	if err != nil {
		return false, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return false, errors.New(message)
		}
		return false, errors.New("non 200 status code")
	}

	selected, err := reply.GetBool("value", true)
	if err != nil {
		return false, err
	}

	return selected, nil

}

func (wd *RemoteWebDriver) IsElementEnabled(element WebElement) (bool, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/enabled", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		nil,
	)

	if err != nil {
		return false, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return false, errors.New(message)
		}
		return false, errors.New("non 200 status code")
	}

	enabled, err := reply.GetBool("value", true)
	if err != nil {
		return false, err
	}

	return enabled, nil

}
func (wd *RemoteWebDriver) GetElementAttribute(element WebElement, name string) (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", wd.URL, wd.Session.GetID(), element.GetWebDriverValue(), name),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	attribute, err := reply.GetString("value", true)
	if err != nil {
		return "", err
	}

	return attribute, nil

}
func (wd *RemoteWebDriver) GetElementProperty(element WebElement, name string) (string, error) {
	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", wd.URL, wd.Session.GetID(), element.GetWebDriverValue(), name),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	property, err := reply.GetString("value", true)
	if err != nil {
		return "", err
	}

	return property, nil
}
func (wd *RemoteWebDriver) GetElementCSS(element WebElement, name string) (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/css/%s", wd.URL, wd.Session.GetID(), element.GetWebDriverValue(), name),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	property, err := reply.GetString("value", true)
	if err != nil {
		return "", err
	}

	return property, nil

}
func (wd *RemoteWebDriver) GetElementText(element WebElement) (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/text", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	property, err := reply.GetString("value", true)
	if err != nil {
		return "", err
	}

	return property, nil

}
func (wd *RemoteWebDriver) GetElementTagName(element WebElement) (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/name", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		nil,
	)

	if err != nil {
		return "", err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return "", errors.New(message)
		}
		return "", errors.New("non 200 status code")
	}

	name, err := reply.GetString("value.qualified name", true)
	if err != nil {
		return "", err
	}

	return name, nil

}
func (wd *RemoteWebDriver) GetElementRect(element WebElement) (*Rect, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/rect", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		nil,
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	x, err := reply.GetFloat("value.x", true)
	if err != nil {
		return nil, err
	}

	y, err := reply.GetFloat("value.y", true)
	if err != nil {
		return nil, err
	}

	width, err := reply.GetFloat("value.width", true)
	if err != nil {
		return nil, err
	}

	height, err := reply.GetFloat("value.height", true)
	if err != nil {
		return nil, err
	}

	return &Rect{X: int(x), Y: int(y), Width: int(width), Height: int(height)}, nil

}
func (wd *RemoteWebDriver) ElementClick(element WebElement) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/click", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		make(map[string]interface{}, 0),
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code received")
	}

	return nil

}
func (wd *RemoteWebDriver) ElementClear(element WebElement) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/clear", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		nil,
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code received")
	}

	return nil

}
func (wd *RemoteWebDriver) ElementSendKeys(element WebElement, keys string) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/value", wd.URL, wd.Session.GetID(), element.GetWebDriverValue()),
		map[string]interface{}{"text": keys},
	)

	if err != nil {
		return err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return errors.New(message)
		}
		return errors.New("non 200 status code received")
	}

	return nil

}

func (wd *RemoteWebDriver) ExecuteScript(script string, args ...interface{}) (interface{}, error) {

	for i := 0; i < len(args); i++ {
		if element, ok := args[i].(WebElement); ok {
			args[i] = map[string]interface{}{element.GetWebDriverID(): element.GetWebDriverValue()}
		}
	}

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/execute/sync", wd.URL, wd.Session.GetID()),
		map[string]interface{}{
			"script": script,
			"args":   args,
		},
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		message, err := reply.GetString("value.message", true)
		if err == nil {
			return nil, errors.New(message)
		}
		return nil, errors.New("non 200 status code received")
	}

	value, err := reply.Get("value", false)
	if err != nil {
		return nil, err
	}

	return value, nil

}
