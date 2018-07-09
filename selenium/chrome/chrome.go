package chrome

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"../../selenium"
)

type chromeDriver struct {
	*selenium.RemoteWebDriver
	cmd *exec.Cmd
}

//Driver starts the chromedriver server on the specified port and returns a WebDriver implementation
func Driver(path string, port int, options *ChromeOptions) (*chromeDriver, error) {

	cmd := exec.Command(path, "--port="+strconv.Itoa(port))

	err := startChromeDriver(cmd)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://127.0.0.1:%d", port)

	caps := &Capabilities{selenium.Capabilities{}, options}

	caps.SetBrowserName("chrome")

	wd := selenium.NewRemote(url, caps)
	if err != nil {
		return nil, err
	}

	driver := &chromeDriver{wd, cmd}

	for retries := 3; retries > 0; retries-- {

		status, err := driver.GetStatus()
		if err != nil {
			if retries == 1 {
				return nil, err
			}
			continue
		}

		if status.IsReady() {
			break
		}

		time.Sleep(5 * time.Second)

	}

	_, err = driver.NewSession()
	if err != nil {
		return nil, err
	}

	return driver, nil

}

func (driver *chromeDriver) NewSession() (*selenium.Session, error) {

	//chrome ,

	reply, err := selenium.ExecuteWDCommand(
		selenium.POST,
		driver.RemoteWebDriver.URL+"/session",
		map[string]interface{}{"desiredCapabilities": driver.Capabilities},
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		return nil, errors.New("non 200 status code received")
	}

	message, err := reply.GetString("value.message", true)
	if err == nil {
		return nil, errors.New(message)
	}

	id, err := reply.GetString("sessionId", false)
	if err != nil {
		return nil, err
	}

	capabilities, err := reply.GetMap("value", false)
	if err != nil {
		return nil, err
	}

	driver.Session = &selenium.Session{ID: id, Capabilities: capabilities}

	return driver.Session, nil

}

func (driver *chromeDriver) GetStatus() (*selenium.Status, error) {

	reply, err := selenium.ExecuteWDCommand(
		selenium.GET,
		fmt.Sprintf("%s/status", driver.URL),
		nil,
	)

	if err != nil {
		return nil, err
	}

	if reply.StatusCode != 200 {
		return nil, errors.New("non 200 status code received")
	}

	message, err := reply.GetString("value.message", true)
	if err == nil {
		return nil, errors.New(message)
	}

	status, err := reply.GetFloat("status", false)
	if err != nil {
		return nil, err
	}

	return &selenium.Status{Ready: status == 0, Message: ""}, nil

}

func (driver *chromeDriver) FindElement(by selenium.By, selector string) (selenium.WebElement, error) {

	element, err := driver.RemoteWebDriver.FindElement(by, selector)
	if err != nil {
		return nil, err
	}

	if webDriverUpdater, ok := element.(selenium.WebDriverUpdater); ok {
		err := webDriverUpdater.SetDriver(driver)
		if err != nil {
			return nil, err
		}
	}

	return element, nil

}

func (driver *chromeDriver) FindElements(by selenium.By, selector string) ([]selenium.WebElement, error) {

	elements, err := driver.RemoteWebDriver.FindElements(by, selector)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		if webDriverUpdater, ok := element.(selenium.WebDriverUpdater); ok {
			err := webDriverUpdater.SetDriver(driver)
			if err != nil {
				return nil, err
			}
		}
	}

	return elements, nil

}

func (driver *chromeDriver) FindElementFromElement(by selenium.By, selector string, element selenium.WebElement) (selenium.WebElement, error) {

	element, err := driver.RemoteWebDriver.FindElementFromElement(by, selector, element)
	if err != nil {
		return nil, err
	}

	if webDriverUpdater, ok := element.(selenium.WebDriverUpdater); ok {
		err := webDriverUpdater.SetDriver(driver)
		if err != nil {
			return nil, err
		}
	}

	return element, nil

}

func (driver *chromeDriver) FindElementsFromElement(by selenium.By, selector string, element selenium.WebElement) ([]selenium.WebElement, error) {

	elements, err := driver.RemoteWebDriver.FindElementsFromElement(by, selector, element)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		if webDriverUpdater, ok := element.(selenium.WebDriverUpdater); ok {
			err := webDriverUpdater.SetDriver(driver)
			if err != nil {
				return nil, err
			}
		}
	}

	return elements, nil

}

/*

func (driver *chromeDriver) GetElementAttribute(element selenium.WebElement, name string) (string, error) {

	reply, err := selenium.ExecuteWDCommand(
		selenium.GET,
		fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", driver.URL, driver.Session.GetID(), element.GetWebDriverValue(), name),
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

	fmt.Println(reply)

	attribute, err := reply.GetString("value.result", true)
	if err != nil {
		return "", err
	}

	return attribute, nil

}

func (driver *chromeDriver) GetElementProperty(element selenium.WebElement, name string) (string, error) {
	reply, err := selenium.ExecuteWDCommand(
		selenium.GET,
		fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", driver.URL, driver.Session.GetID(), element.GetWebDriverID(), name),
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

	property, err := reply.GetString("value.result", true)
	if err != nil {
		return "", err
	}

	return property, nil
}

*/

func (driver *chromeDriver) ElementSendKeys(element selenium.WebElement, keys string) error {

	chars := make([]string, len(keys))
	for i, c := range keys {
		chars[i] = string(c)
	}

	reply, err := selenium.ExecuteWDCommand(
		selenium.POST,
		fmt.Sprintf("%s/session/%s/element/%s/value", driver.URL, driver.Session.GetID(), element.GetWebDriverValue()),
		map[string]interface{}{"value": chars},
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

func (driver *chromeDriver) Quit() error {

	err := driver.DeleteSession()
	if err != nil {
		return err
	}

	return driver.cmd.Process.Kill()

}

func (driver *chromeDriver) SetTimeouts(timeouts *selenium.Timeouts) error {

	list := []map[string]interface{}{

		map[string]interface{}{
			"type": "page load",
			"ms":   timeouts.PageLoad,
		}, map[string]interface{}{
			"type": "script",
			"ms":   timeouts.Script,
		},
		map[string]interface{}{
			"type": "implicit",
			"ms":   timeouts.Implicit,
		},
	}

	for _, timeout := range list {

		reply, err := selenium.ExecuteWDCommand(
			selenium.POST,
			fmt.Sprintf("%s/session/%s/timeouts", driver.URL, driver.Session.GetID()),
			timeout,
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

	}

	return nil

}

func (driver *chromeDriver) GetElementRect(element selenium.WebElement) (*selenium.Rect, error) {

	returned, err := driver.RemoteWebDriver.ExecuteScript("return arguments[0].getBoundingClientRect()", element)
	if err != nil {
		return nil, err
	}

	if data, ok := returned.(map[string]interface{}); ok {

		x, ok := data["x"].(float64)
		if !ok {
			return nil, errors.New("'x' not found")
		}

		y, ok := data["y"].(float64)
		if !ok {
			return nil, errors.New("'y' not found")
		}

		width, ok := data["width"].(float64)
		if !ok {
			return nil, errors.New("'width' not found")
		}

		height, ok := data["height"].(float64)
		if !ok {
			return nil, errors.New("'height' not found")
		}

		return &selenium.Rect{X: int(x), Y: int(y), Width: int(width), Height: int(height)}, nil

	}

	return nil, errors.New("could not parse rect")

}

func startChromeDriver(cmd *exec.Cmd) error {

	/*

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return err
		}

	*/

	if err := cmd.Start(); err != nil {
		return err
	}

	//go func() { selenium.Log(stdout) }()
	//go func() { selenium.Log(stderr) }()

	return nil
}
