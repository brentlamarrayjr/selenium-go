package by

type Locator struct {
	By       By
	Location string
}

func NewLocator(by By, location string) (*Locator, error) {

	return &Locator{By: by, Location: location}, nil

}
