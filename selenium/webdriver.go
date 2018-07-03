package selenium

type WebDriver interface {
	NewSession() (Session, error)
	GetTimeouts() (*Timeouts, error)
	SetTimeouts(*Timeouts) error
	GetStatus() (status, error)
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
	GetWindowRect() (map[string]interface{}, error)
	SetWindowRect(width int, height int, x int, y int) error
	MaximizeWindow() error
	MinimizeWindow() error
	FullscreenWindow() error
	FindElement(by By, selector string) (WebElement, error)
	FindElements(by By, selector string) ([]WebElement, error)
	FindElementFromElement(by By, selector string, id string) (WebElement, error)
	FindElementsFromElement(by By, selector string, id string) ([]WebElement, error)
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
}
