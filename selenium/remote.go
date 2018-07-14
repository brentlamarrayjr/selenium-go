package selenium

import (
	"errors"
	"fmt"

	"./by"
)

type remoteWebDriver struct {
	session             *session
	desiredCapabilities Capabilities `json:"capabilities,omitempty"`
	url                 string
}

//NewRemote returns a pointer to an implementation of the W3C WebDriver client protocol
func NewRemote(url string, desiredCapabilities Capabilities) *remoteWebDriver {

	return &remoteWebDriver{desiredCapabilities: desiredCapabilities, url: url}

}

func (wd *remoteWebDriver) SetSession(id string, caps map[string]interface{}) {

	wd.session = &session{id, caps}
}

func (wd *remoteWebDriver) GetURL() string {
	return wd.url
}

func (wd *remoteWebDriver) GetDesiredCapabilities() Capabilities {
	return wd.desiredCapabilities
}

func (wd *remoteWebDriver) GetSession() SessionInfo {
	return wd.session
}

//NewSession creates a single instantiation of a particular user agent and returns the session ID.
func (wd *remoteWebDriver) NewSession() (SessionInfo, error) {

	//chrome map[string]interface{}{"desiredCapabilities": wd.Capabilities},

	reply, err := ExecuteWDCommand(
		POST,
		wd.url+"/session",
		map[string]interface{}{
			"alwaysMatch": wd.desiredCapabilities,
		},
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(reply.Data)

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

	wd.session = &session{ID: id, Capabilities: capabilities}

	return wd.session, nil

}

func (wd *remoteWebDriver) GetStatus() (*Status, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/status", wd.url),
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

func (wd *remoteWebDriver) GetTimeouts() (*Timeouts, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/timeouts", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) SetTimeouts(timeouts *Timeouts) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/timeouts", wd.url, wd.session.GetID()),
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
func (wd *remoteWebDriver) DeleteSession() error {

	reply, err := ExecuteWDCommand(
		DELETE,
		fmt.Sprintf("%s/session/%s", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) Navigate(url string) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/url", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) GetCurrentURL() (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/url", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) Back() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/back", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) Forward() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/forward", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) Refresh() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/refresh", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) GetTitle() (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/title", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) GetWindowHandle() (string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/window", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) CloseWindow() error {

	reply, err := ExecuteWDCommand(
		DELETE,
		fmt.Sprintf("%s/session/%s/window", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) SwitchToWindow(window string) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) GetWindowHandles() ([]string, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/window/handles", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) SwitchToFrame(id int) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/frame", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) SwitchToParentFrame() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/frame/parent", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) GetWindowRect() (*Rect, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/window/rect", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) SetWindowRect(rect *Rect) error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/rect", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) MaximizeWindow() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/maximize", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) MinimizeWindow() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/minimize", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) FullscreenWindow() error {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/window/fullscreen", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) FindElement(locator *by.Locator) (WebElement, error) {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element", wd.url, wd.session.GetID()),
		map[string]interface{}{
			"using": locator.By,
			"value": locator.Location,
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

func (wd *remoteWebDriver) FindElements(locator *by.Locator) ([]WebElement, error) {

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element", wd.url, wd.session.GetID()),
		map[string]interface{}{
			"using": locator.By,
			"value": locator.Location,
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

func (wd *remoteWebDriver) FindElementFromElement(element WebElement, locator *by.Locator) (WebElement, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return nil, errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/element", wd.url, wd.session.GetID(), info.GetValue()),
		map[string]interface{}{
			"using": locator.By,
			"value": locator.Location,
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

func (wd *remoteWebDriver) FindElementsFromElement(element WebElement, locator *by.Locator) ([]WebElement, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return nil, errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/element", wd.url, wd.session.GetID(), info.GetValue()),
		map[string]interface{}{
			"using": locator.By,
			"value": locator.Location,
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

func (wd *remoteWebDriver) GetActiveElement() (WebElement, error) {

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/active", wd.url, wd.session.GetID()),
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

func (wd *remoteWebDriver) IsElementSelected(element WebElement) (bool, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return false, errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/selected", wd.url, wd.session.GetID(), info.GetValue()),
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

func (wd *remoteWebDriver) IsElementEnabled(element WebElement) (bool, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return false, errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/enabled", wd.url, wd.session.GetID(), info.GetValue()),
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
func (wd *remoteWebDriver) GetElementAttribute(element WebElement, name string) (string, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return "", errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", wd.url, wd.session.GetID(), info.GetValue(), name),
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
func (wd *remoteWebDriver) GetElementProperty(element WebElement, name string) (string, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return "", errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", wd.url, wd.session.GetID(), info.GetValue(), name),
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
func (wd *remoteWebDriver) GetElementCSS(element WebElement, name string) (string, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return "", errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/css/%s", wd.url, wd.session.GetID(), info.GetValue(), name),
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
func (wd *remoteWebDriver) GetElementText(element WebElement) (string, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return "", errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/text", wd.url, wd.session.GetID(), info.GetValue()),
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
func (wd *remoteWebDriver) GetElementTagName(element WebElement) (string, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return "", errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/name", wd.url, wd.session.GetID(), info.GetValue()),
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
func (wd *remoteWebDriver) GetElementRect(element WebElement) (*Rect, error) {

	info, ok := element.(WebElementInfo)
	if !ok {
		return nil, errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		GET,
		fmt.Sprintf("%s/session/%s/element/%s/rect", wd.url, wd.session.GetID(), info.GetValue()),
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
func (wd *remoteWebDriver) ElementClick(element WebElement) error {

	info, ok := element.(WebElementInfo)
	if !ok {
		return errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/click", wd.url, wd.session.GetID(), info.GetValue()),
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
func (wd *remoteWebDriver) ElementClear(element WebElement) error {

	info, ok := element.(WebElementInfo)
	if !ok {
		return errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/clear", wd.url, wd.session.GetID(), info.GetValue()),
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
func (wd *remoteWebDriver) ElementSendKeys(element WebElement, keys string) error {

	info, ok := element.(WebElementInfo)
	if !ok {
		return errors.New("could not get web element info")
	}

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/element/%s/value", wd.url, wd.session.GetID(), info.GetValue()),
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

func (wd *remoteWebDriver) ExecuteScript(script string, args ...interface{}) (interface{}, error) {

	for i := 0; i < len(args); i++ {
		if element, ok := args[i].(WebElement); ok {

			info, ok := element.(WebElementInfo)
			if !ok {
				return nil, errors.New("could not get web element info")
			}

			args[i] = map[string]interface{}{info.GetID(): info.GetValue()}
		}
	}

	reply, err := ExecuteWDCommand(
		POST,
		fmt.Sprintf("%s/session/%s/execute/sync", wd.url, wd.session.GetID()),
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
