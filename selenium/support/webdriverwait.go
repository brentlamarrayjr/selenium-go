package support

import (
	"errors"

	"../../selenium"
)

type webDriverWait struct {
	driver  selenium.WebDriver
	timeout int
}

func WebDriverWait(driver selenium.WebDriver, timeout int) *webDriverWait {

	return &webDriverWait{driver: driver, timeout: timeout}

}

func (wait *webDriverWait) Until(ec ExpectedCondition) error {
	for i := 0; i < wait.timeout; i++ {
		if ec.Wait(wait.driver) == nil {
			return nil
		}
	}
	return errors.New("Wait timeout")
}
