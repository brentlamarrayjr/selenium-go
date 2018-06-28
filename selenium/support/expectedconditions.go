package support

import (
	"errors"

	"../../selenium"
)

type PresenceOfElementLocated struct{}
type ElementToBeDisplayed struct{}

func (ec *ElementToBeDisplayed) Wait(selenium.WebDriver) error {

	

	return errors.New("")
}

func (ec *PresenceOfElementLocated) Wait(selenium.WebDriver) error {

	return errors.New("")
}
