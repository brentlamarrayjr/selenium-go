package selenium

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
func (e *webElement) GetWebDriverID() string { return e.id }

/* Get WebDriver ID */
func (e *webElement) GetWebDriverValue() string { return e.value }

/* Click on element */
func (e *webElement) Click() error { return e.driver.ElementClick(e) }

/* Send keys (type) into element */
func (e *webElement) SendKeys(keys string) error { return e.driver.ElementSendKeys(e, keys) }

/* Submit */
func (e *webElement) Submit() error { return nil }

/* Clear */
func (e *webElement) Clear() error { return e.driver.ElementClear(e) }

/* Find children, return one element. */
func (e *webElement) FindElement(by By, selection string) (WebElement, error) {
	return e.driver.FindElementFromElement(by, selection, e)
}

/* Find children, return list of elements. */
func (e *webElement) FindElements(by By, selection string) ([]WebElement, error) {
	return e.driver.FindElementsFromElement(by, selection, e)
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

	/* Get WebDriver ID */
	GetWebDriverID() string
	/* Get WebDriver Value */
	GetWebDriverValue() string

	/* Click on element */
	Click() error
	/* Send keys (type) into element */
	SendKeys(keys string) error
	/* Submit */
	Submit() error
	/* Clear */
	Clear() error
	/* Find children, return one element. */
	FindElement(by By, selection string) (WebElement, error)
	/* Find children, return list of elements. */
	FindElements(by By, selection string) ([]WebElement, error)
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
