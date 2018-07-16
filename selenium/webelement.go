package selenium

import "./by"

type webElement struct {
	id     string
	value  string
	driver WebDriver
}

func (e *webElement) SetDriver(driver WebDriver) error {

	e.driver = driver
	return nil

	//return errors.New("could not cast 'driver' to WebDriber")

}

/* Get WebDriver ID */
func (e *webElement) GetID() string { return e.id }

/* Get WebDriver ID */
func (e *webElement) GetValue() string { return e.value }

/* Click on element */
func (e *webElement) Click() error { return e.driver.ElementClick(e) }

/* Send keys (type) into element */
func (e *webElement) SendKeys(keys string) error { return e.driver.ElementSendKeys(e, keys) }

/* Submit performs the submit action on a form or form control */
func (e *webElement) Submit() error {

	form, err := e.FindElement(by.XPath("./ancestor-or-self::form"))
	if err != nil {
		return err
	}

	_, err = e.driver.ExecuteScript("arguments[0].submit()", form)

	//return e.driver.ElementSendKeys(e, string(keys.Enter))
	return err

}

/* Clear clears an input element */
func (e *webElement) Clear() error { return e.driver.ElementClear(e) }

/* FindElement returns one WebElement found via Locator. */
func (e *webElement) FindElement(locator *by.Locator) (WebElement, error) {
	return e.driver.FindElementFromElement(e, locator)
}

/* FindElements return list of elements found via Locator. */
func (e *webElement) FindElements(locator *by.Locator) ([]WebElement, error) {
	return e.driver.FindElementsFromElement(e, locator)
}

/* GetTagName returns the WebElement tag name */
func (e *webElement) GetTagName() (string, error) { return e.driver.GetElementTagName(e) }

/*GetText return the text of a WebElement */
func (e *webElement) GetText() (string, error) { return e.driver.GetElementText(e) }

/*IsSelected return a boolean that indicates if the WebElement is selected. */
func (e *webElement) IsSelected() (bool, error) { return e.driver.IsElementSelected(e) }

/*IsEnabled return a boolean that indicates if the WebElement is enabled. */
func (e *webElement) IsEnabled() (bool, error) { return e.driver.IsElementEnabled(e) }

/*IsDisplayed return a boolean that indicates if the WebElement is selected. */
func (e *webElement) IsDisplayed() (bool, error) {

	elementRect, err := e.driver.GetElementRect(e)
	if err != nil {
		return false, err
	}
	windowRect, err := e.driver.GetWindowRect()
	if err != nil {
		return false, err
	}

	return elementRect.Overlaps(windowRect), nil

}

/* Get element attribute. */
func (e *webElement) GetAttribute(name string) (string, error) {
	return e.driver.GetElementAttribute(e, name)
}

/* Get element property. */
func (e *webElement) GetProperty(name string) (string, error) {
	return e.driver.GetElementProperty(e, name)
}

/* Element location: x, y.*/
func (e *webElement) GetRect() (*Rect, error) { return e.driver.GetElementRect(e) }

/* Get element CSS property value. */
func (e *webElement) GetCSS(name string) (string, error) { return e.driver.GetElementCSS(e, name) }

//WebElement provides an interface to common actions performed on a Selenium WebElement
type WebElement interface {
	Click() error
	SendKeys(keys string) error
	Submit() error
	Clear() error
	FindElement(locator *by.Locator) (WebElement, error)
	FindElements(locator *by.Locator) ([]WebElement, error)
	GetTagName() (string, error)
	GetText() (string, error)
	IsSelected() (bool, error)
	IsEnabled() (bool, error)
	IsDisplayed() (bool, error)
	GetAttribute(name string) (string, error)
	GetProperty(name string) (string, error)
	GetRect() (*Rect, error)
	GetCSS(name string) (string, error)
}
