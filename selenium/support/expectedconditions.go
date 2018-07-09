package support

import (
	"errors"

	"../../selenium"
)

type expectedCondition struct {
	wait func(selenium.WebDriver) error
}

func (ec expectedCondition) Wait(driver selenium.WebDriver) error {
	return ec.wait(driver)
}

type elementToBeDisplayed struct{}

func PresenceOfElementLocated(locator *Locator) ExpectedCondition {

	return &expectedCondition{
		wait: func(driver selenium.WebDriver) error {
			_, err := driver.FindElement(locator.By, locator.Location)
			if err != nil {
				return err
			}
			return nil
		},
	}

}

func ElementToBeDisplayed(locator *Locator) ExpectedCondition {

	return &expectedCondition{
		wait: func(driver selenium.WebDriver) error {
			element, err := driver.FindElement(locator.By, locator.Location)
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
