package selenium

type WebDriver interface {
	NewSession() (session, error)
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
	FindElement(by By, selector string) (map[string]interface{}, error)
	FindElements(by By, selector string) ([]map[string]interface{}, error)
	FindElementFromElement(by By, selector string, id string) (map[string]interface{}, error)
	FindElementsFromElement(by By, selector string, id string) ([]map[string]interface{}, error)
	GetActiveElement() (map[string]interface{}, error)
	IsElementSelected(by By, selector string, id string) ([]map[string]interface{}, error)
}
