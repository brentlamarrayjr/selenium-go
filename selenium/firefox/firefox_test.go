package firefox

import (
	"fmt"
	"testing"

	"../../selenium"
	"github.com/stretchr/testify/require"
)

func TestFirefox(t *testing.T) {

	//options := new(ChromeOptions)
	driver, err := Driver("C:\\Users\\brent\\Development\\Go\\Projects\\selenium\\selenium\\firefox\\geckodriver.exe", 4444, nil)
	require.NoErrorf(t, err, "Creation of GeckoDriver should not raise any errors.")
	err = driver.Navigate("https://google.com")
	require.NoErrorf(t, err, "GeckoDriver navigation should not raise any errors.")
	element, err := driver.FindElement(selenium.XPath, "//input[@name='q']")
	require.NoErrorf(t, err, "GeckoDriver find element should not raise any errors.")
	fmt.Println(element)
	require.NoErrorf(t, driver.Quit(), "Quitting of GeckoDriver should not raise any errors.")

}
