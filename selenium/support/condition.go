package support

import "../../selenium"

type ExpectedCondition interface {
	Wait(selenium.WebDriver) error
}
