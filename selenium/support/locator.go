package support

import "../../selenium"

type locator struct {
	by       selenium.By
	location string
}

func Locator(by selenium.By, location string) (*locator, error) {

	return &locator{by: by, location: location}, nil

}
