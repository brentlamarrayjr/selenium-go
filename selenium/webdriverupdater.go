package selenium

type WebDriverUpdater interface {
	SetDriver(driver WebDriver) error
}
