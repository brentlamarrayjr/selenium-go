package selenium

type WebElementUpdater interface {
	SetDriver(driver WebDriver) error
}
