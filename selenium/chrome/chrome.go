package chrome

import (
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

func startChromeDriver(cmd *exec.Cmd) error {

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	go func() { selenium.Log(stdout) }()
	go func() { selenium.Log(stderr) }()

	return nil
}
