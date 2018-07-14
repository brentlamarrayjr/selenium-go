package selenium

type WebDriverUpdater interface {
	SetSession(id string, caps map[string]interface{})
}
