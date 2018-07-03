package selenium

type element struct {
	ID string `json:ELEMENT`
}

type webElement struct {
	id     string
	driver WebDriver
}

/* Get WebDriver ID */
func (e *webElement) GetWebDriverID() string { return "" }

/* Click on element */
func (e *webElement) Click() error { return nil }

/* Send keys (type) into element */
func (e *webElement) SendKeys(keys string) error { return nil }

/* Submit */
func (e *webElement) Submit() error { return nil }

/* Clear */
func (e *webElement) Clear() error { return nil }

/* Find children, return one element. */
func (e *webElement) FindElement(by, value string) (WebElement, error) { return nil, nil }

/* Find children, return list of elements. */
func (e *webElement) FindElements(by, value string) ([]WebElement, error) { return nil, nil }

/* Element name */
func (e *webElement) GetTagName() (string, error) { return "", nil }

/* Text of element */
func (e *webElement) GetText() (string, error) { return "", nil }

/* Check if element is selected. */
func (e *webElement) IsSelected() (bool, error) { return false, nil }

/* Check if element is enabled. */
func (e *webElement) IsEnabled() (bool, error) { return false, nil }

/* Check if element is displayed. */
func (e *webElement) IsDisplayed() (bool, error) { return false, nil }

/* Get element attribute. */
func (e *webElement) GetAttribute(name string) (string, error) { return "", nil }

/* Get element property. */
func (e *webElement) GetProperty(name string) (string, error) { return "", nil }

/* Element location: x, y.*/
func (e *webElement) GetLocation() (int, int, error) { return 0, 0, nil }

/* Element size: width, height */
func (e *webElement) GetSize() (int, int, error) { return 0, 0, nil }

/* Get element CSS property value. */
func (e *webElement) GetCSS(name string) (string, error) { return "", nil }

type WebElement interface {

	/* Get WebDriver ID */
	GetWebDriverID() string

	/* Click on element */
	Click() error
	/* Send keys (type) into element */
	SendKeys(keys string) error
	/* Submit */
	Submit() error
	/* Clear */
	Clear() error

	/* Find children, return one element. */
	FindElement(by, value string) (WebElement, error)
	/* Find children, return list of elements. */
	FindElements(by, value string) ([]WebElement, error)

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

	/* Element location: x, y.*/
	GetLocation() (int, int, error)
	/* Element size: width, height */
	GetSize() (int, int, error)
	/* Get element CSS property value. */
	GetCSS(name string) (string, error)
}
