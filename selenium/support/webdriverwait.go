package support

import (
	"errors"
	"time"

	"../../selenium"
)

type webDriverWait struct {
	driver  selenium.WebDriver
	timeout int
	poll    int
}

func WebDriverWait(driver selenium.WebDriver, timeout int, poll int) *webDriverWait {

	return &webDriverWait{driver: driver, timeout: timeout, poll: poll}

}

func (wait *webDriverWait) Until(ec ExpectedCondition) error {

	timeoutTicker := time.NewTicker(time.Duration(wait.timeout) * time.Second)
	pollTicker := time.NewTicker(time.Duration(wait.poll) * time.Second)
	conditionChan := make(chan bool)
	errChan := make(chan error)

	go func() {
		for range timeoutTicker.C {
			errChan <- errors.New("Timeout error")
		}
	}()
	go func() {
		for range pollTicker.C {
			if ec.Wait(wait.driver) == nil {
				conditionChan <- true
			}
			conditionChan <- false
		}
	}()

	for {
		select {
		case err := <-errChan:
			close(conditionChan)
			close(errChan)
			return err
		case condition := <-conditionChan:
			if condition {
				close(errChan)
				close(conditionChan)
				return nil
			}
		}
	}

}
