package selenium

import "./by"

type WebDriver interface {
	NewSession() (*Session, error)
	GetTimeouts() (*Timeouts, error)
	SetTimeouts(*Timeouts) error
	GetStatus() (*Status, error)
	DeleteSession() error
	Navigate(url string) error
	GetCurrentURL() (string, error)
	Back() error
	Forward() error
	Refresh() error
	GetTitle() (string, error)
	GetWindowHandle() (string, error)
	CloseWindow() error
	SwitchToWindow(window string) error
	GetWindowHandles() ([]string, error)
	SwitchToFrame(id int) error
	SwitchToParentFrame() error
	GetWindowRect() (*Rect, error)
	SetWindowRect(rect *Rect) error
	MaximizeWindow() error
	MinimizeWindow() error
	FullscreenWindow() error
	FindElement(locator *by.Locator) (WebElement, error)
	FindElements(locator *by.Locator) ([]WebElement, error)
	FindElementFromElement(element WebElement, locator *by.Locator) (WebElement, error)
	FindElementsFromElement(element WebElement, locator *by.Locator) ([]WebElement, error)
	GetActiveElement() (WebElement, error)
	IsElementSelected(element WebElement) (bool, error)
	IsElementEnabled(element WebElement) (bool, error)
	GetElementAttribute(element WebElement, name string) (string, error)
	GetElementProperty(element WebElement, name string) (string, error)
	GetElementCSS(element WebElement, name string) (string, error)
	GetElementText(element WebElement) (string, error)
	GetElementTagName(element WebElement) (string, error)
	GetElementRect(element WebElement) (*Rect, error)
	ElementClick(element WebElement) error
	ElementClear(element WebElement) error
	ElementSendKeys(element WebElement, keys string) error
	ExecuteScript(script string, args ...interface{}) (interface{}, error)
}
