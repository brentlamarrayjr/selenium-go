package support

import (
	"errors"
	"sync"
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

	errChan := make(chan error)

	mutex := &sync.Mutex{}
	closed := false

	//Start listening for timeout
	go func() {
		for range timeoutTicker.C {
			mutex.Lock()
			if !closed {
				errChan <- errors.New("timeout error")
				timeoutTicker.Stop()
				pollTicker.Stop()
				close(errChan)
			}
			mutex.Unlock()
		}
	}()

	//Start polling for condition
	go func() {
		for range pollTicker.C {
			mutex.Lock()
			if !closed {
				if ec.Wait(wait.driver) == nil {
					closed = true
					errChan <- nil
					timeoutTicker.Stop()
					pollTicker.Stop()
					close(errChan)
				}
			}
			mutex.Unlock()
		}
	}()

	return <-errChan

}
