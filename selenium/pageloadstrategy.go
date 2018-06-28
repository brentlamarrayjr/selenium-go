package selenium

type PageLoadStrategy string

const (
	None   PageLoadStrategy = "none"
	Eager  PageLoadStrategy = "eager"
	Normal PageLoadStrategy = "normal"
)
