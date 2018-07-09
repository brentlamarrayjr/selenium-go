package selenium

import "errors"

type PageLoadStrategy string

const (
	None   PageLoadStrategy = "none"
	Eager  PageLoadStrategy = "eager"
	Normal PageLoadStrategy = "normal"
)

func ParsePageLoadStrategy(strategy string) (PageLoadStrategy, error) {
	switch strategy {

	case string(Normal):
		return Normal, nil
	case string(Eager):
		return Eager, nil
	case string(None):
		return None, nil

	}

	return Normal, errors.New("invalid page load strategy")

}
