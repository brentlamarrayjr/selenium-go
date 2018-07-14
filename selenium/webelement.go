package selenium

import "./by"

type webElement struct {
	id     string
	value  string
	driver WebDriver
}

func (element *webElement) SetDriver(driver WebDriver) error {

	element.driver = driver
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

/* Submit */
func (e *webElement) Submit() error {

	form, err := e.FindElement(by.XPath("./ancestor-or-self::form"))
	if err != nil {
		return err
	}

	_, err = e.driver.ExecuteScript("arguments[0].submit()", form)

	//return e.driver.ElementSendKeys(e, string(keys.Enter))
	return err

}

/* Clear */
func (e *webElement) Clear() error { return e.driver.ElementClear(e) }

/* Find children, return one element. */
func (e *webElement) FindElement(locator *by.Locator) (WebElement, error) {
	return e.driver.FindElementFromElement(e, locator)
}

/* Find children, return list of elements. */
func (e *webElement) FindElements(locator *by.Locator) ([]WebElement, error) {
	return e.driver.FindElementsFromElement(e, locator)
}

/* Element name */
func (e *webElement) GetTagName() (string, error) { return e.driver.GetElementTagName(e) }

/* Text of element */
func (e *webElement) GetText() (string, error) { return e.driver.GetElementText(e) }

/* Check if element is selected. */
func (e *webElement) IsSelected() (bool, error) { return e.driver.IsElementSelected(e) }

/* Check if element is enabled. */
func (e *webElement) IsEnabled() (bool, error) { return e.driver.IsElementEnabled(e) }

/* Check if element is displayed. */
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

type WebElement interface {

	/* Click on element */
	Click() error
	/* Send keys (type) into element */
	SendKeys(keys string) error
	/* Submit */
	Submit() error
	/* Clear */
	Clear() error
	/* Find children, return one element. */
	FindElement(locator *by.Locator) (WebElement, error)
	/* Find children, return list of elements. */
	FindElements(locator *by.Locator) ([]WebElement, error)
	/* Element name */
	GetTagName() (string, error)
	/* Text of element */
	GetText() (string, error)
	/* Check if element is selected. */
	IsSelected() (bool, error)
	/* Check if element is enabled. */
	IsEnabled() (bool, error)
	/* Check if element is displayed. */
	IsDisplayed() (bool, error)
	/* Get element attribute. */
	GetAttribute(name string) (string, error)
	/* Get element property. */
	GetProperty(name string) (string, error)
	/* Element rect: x, y, height, width.*/
	GetRect() (*Rect, error)
	/* Get element CSS property value. */
	GetCSS(name string) (string, error)
}
