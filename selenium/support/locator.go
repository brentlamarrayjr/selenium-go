package support

import "../../selenium"

type Locator struct {
	By       selenium.By
	Location string
}

func NewLocator(by selenium.By, location string) (*Locator, error) {

	return &Locator{By: by, Location: location}, nil

}
