package chrome

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

	driver := selenium.NewRemote(url, caps)
	if err != nil {
		return nil, err
	}

	for retries := 3; retries > 0; retries-- {

		time.Sleep(5 * time.Second)

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

	}

	session, err := driver.NewSession()
	if err != nil {
		return nil, err
	}

	fmt.Println("Session:")
	fmt.Println(session)
	fmt.Println(session.GetReturnedCapabilities())

	return &chromeDriver{driver, cmd}, nil

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

		jsonData, err := json.Marshal(timeout)
		if err != nil {
			return err
		}

		fmt.Printf("%s/session/%s/timeouts\n", driver.URL, driver.Session.GetID())

		resp, err := http.Post(fmt.Sprintf("%s/session/%s/timeouts", driver.URL, driver.Session.GetID()), "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		reply := new(selenium.Reply)
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

	}

	return nil

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
