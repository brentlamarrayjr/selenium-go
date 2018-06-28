package selenium

type By string

const (
	CSS             By = "css selector"
	LinkText        By = "link text"
	PartialLinkText By = "partial link text"
	TagName         By = "tag name"
	XPath           By = "xpath"
)
