package firefox

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"../../selenium"
)

type Options struct {
	Args    []string               `json:",omitempty"`
	Binary  string                 `json:",omitempty"`
	Profile string                 `json:",omitempty"`
	Log     Log                    `json:",omitempty"`
	Prefs   map[string]interface{} `json:",omitempty"`
}

func (options *Options) AddArgs(args ...string) {
	options.Args = append(options.Args, args...)
}

func (options *Options) AddPref(name string, value interface{}) {
	options.Prefs[name] = value
}

//Log sets the logging verbosity of geckodriver and Firefox, including all specified level logs and above.
type Log struct {
	Level string `json:"level,omitempty"`
}

type geckoDriver struct {
	*selenium.RemoteWebDriver
	cmd *exec.Cmd
}

//GeckoDriver starts the geckodriver server on the specified port and returns a WebDriver implementation
func Driver(path string, port int, options *Options) (*geckoDriver, error) {

	cmd := exec.Command(path, "--port="+strconv.Itoa(port))
	err := startGeckoDriver(cmd)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://127.0.0.1:%d", port)

	caps := &Capabilities{selenium.Capabilities{}, options}
	caps.SetBrowserName("firefox")

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

	return &geckoDriver{driver, cmd}, nil

}

func startGeckoDriver(cmd *exec.Cmd) error {

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

func (driver *geckoDriver) Quit() error {

	err := driver.DeleteSession()
	if err != nil {
		return err
	}

	/*

		i, err := strconv.Atoi(driver.RemoteWebDriver.Session.GetReturnedCapabilities().ProcessID)
		if err != nil {
			return err
		}



			process, err := os.FindProcess(i)
			if err != nil {
				return err
			}

	*/

	return driver.cmd.Process.Kill()

}
