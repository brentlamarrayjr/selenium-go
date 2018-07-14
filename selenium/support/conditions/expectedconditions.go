package conditions

import (
	"errors"

	"../../../selenium"
	"../../by"
	"../../support"
)

type expectedCondition struct {
	wait func(selenium.WebDriver) error
}

func (ec expectedCondition) Wait(driver selenium.WebDriver) error {
	return ec.wait(driver)
}

func PresenceOfElementLocated(locator *by.Locator) support.ExpectedCondition {

	return &expectedCondition{
		wait: func(driver selenium.WebDriver) error {
			_, err := driver.FindElement(locator)
			if err != nil {
				return err
			}
			return nil
		},
	}

}

func ElementToBeDisplayed(locator *by.Locator) support.ExpectedCondition {

	return &expectedCondition{
		wait: func(driver selenium.WebDriver) error {
			element, err := driver.FindElement(locator)
			if err != nil {
				return err
			}

			displayed, err := element.IsDisplayed()
			if err != nil {
				return err
			}

			if !displayed {
				return errors.New("Not found")
			}

			return nil
		},
	}

}
