package by

import (
	"fmt"
)

type By string

const (
	css             By = "css selector"
	linkText        By = "link text"
	partialLinkText By = "partial link text"
	tagName         By = "tag name"
	xpath           By = "xpath"
)

func ID(id string) *Locator {
	return &Locator{css, fmt.Sprintf("#%s", id)}
}

func Name(name string) *Locator {
	return &Locator{css, fmt.Sprintf("input[name='%s']", name)}
}

func CSS(selector string) *Locator {
	return &Locator{css, selector}
}

func Tag(name string) *Locator {
	return &Locator{tagName, name}
}

func LinkText(text string) *Locator {
	return &Locator{linkText, text}
}

func PartialLinkText(text string) *Locator {
	return &Locator{partialLinkText, text}
}

func XPath(selector string) *Locator {
	return &Locator{xpath, selector}
}
