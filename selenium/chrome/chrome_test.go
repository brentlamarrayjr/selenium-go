package chrome

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"../../selenium/by"
	"../../selenium/support"
	"../../selenium/support/conditions"
	"github.com/stretchr/testify/require"
)

func TestChrome(t *testing.T) {

	//options := new(ChromeOptions)
	driver, err := Driver(CHROMEDRIVER_PATH, 4444, nil)
	require.NoErrorf(t, err, "Creation of ChromeDriver should not raise any errors.")

	timeouts, err := driver.GetTimeouts()
	require.NoErrorf(t, err, "ChromeDriver getting timeouts should not raise any errors.")

	timeouts.Implicit = 5000
	err = driver.SetTimeouts(timeouts)
	require.NoErrorf(t, err, "ChromeDriver setting timeouts should not raise any errors.")

	timeouts2, err := driver.GetTimeouts()
	require.NoErrorf(t, err, "ChromeDriver getting timeouts should not raise any errors.")
	require.Equalf(t, timeouts, timeouts2, "ChromeDriver getting timeouts should not raise any errors.")

	err = driver.Navigate("https://google.com")
	require.NoErrorf(t, err, "ChromeDriver navigation should not raise any errors.")

	time.Sleep(3 * time.Second)

	element, err := driver.FindElement(by.XPath("//input[@name='q']"))
	require.NoErrorf(t, err, "ChromeDriver find element should not raise any errors.")
	fmt.Println(fmt.Sprintf("Element: %v", element))

	name, err := element.GetAttribute("name")
	require.NoErrorf(t, err, "ChromeDriver send keys should not raise any errors.")
	require.Equalf(t, name, "q", "ChromeDriver get element attribute should not raise any errors.")
	fmt.Println(fmt.Sprintf("Element name: %s", name))

	rect, err := element.GetRect()
	require.NoErrorf(t, err, "ChromeDriver send keys should not raise any errors.")
	require.Equalf(t, name, "q", "ChromeDriver get element attribute should not raise any errors.")
	fmt.Println(fmt.Sprintf("Rect: %v", rect))

	isDisplayed, err := element.IsDisplayed()
	require.NoErrorf(t, err, "ChromeDriver send keys should not raise any errors.")
	require.Truef(t, isDisplayed, "ChromeDriver is element displayed should not raise any errors.")
	fmt.Println(fmt.Sprintf("Element is displayed: %t", isDisplayed))

	isEnabled, err := element.IsEnabled()
	require.NoErrorf(t, err, "ChromeDriver send keys should not raise any errors.")
	require.Truef(t, isEnabled, "ChromeDriver is element enabled should not raise any errors.")
	fmt.Println(fmt.Sprintf("Element is enabled: %t", isEnabled))

	err = element.SendKeys("automation testing")
	require.NoErrorf(t, err, "ChromeDriver send keys should not raise any errors.")

	err = element.Submit()
	require.NoErrorf(t, err, "ChromeDriver submit should not raise any errors.")

	//searchButton, err := driver.FindElement(by.XPath("//input[@value='Google Search']"))
	//require.NoErrorf(t, err, "ChromeDriver find element should not raise any errors.")

	//err = searchButton.Click()
	//require.NoErrorf(t, err, "ChromeDriver click element should not raise any errors.")

	locator := by.XPath("//div[@id='resultStats']")
	require.NoErrorf(t, err, "Locator creation should not raise any errors.")

	err = support.WebDriverWait(driver, 10, 1).Until(
		conditions.PresenceOfElementLocated(locator),
	)

	require.NoErrorf(t, err, "ChromeDriver WebDriverWait should not raise any errors.")

	searchInfo, err := driver.FindElement(by.XPath("//div[@id='resultStats']"))
	require.NoErrorf(t, err, "ChromeDriver find element should not raise any errors.")

	text, err := searchInfo.GetText()
	require.NoErrorf(t, err, "ChromeDriver get element text should not raise any errors.")

	require.Regexpf(
		t,
		"About \\d{1,3}(,\\d{3})*(\\.\\d+)? results \\(\\d*\\.?\\d+ seconds\\)",
		strings.Trim(text, " "),
		"Element text should match expected",
	)

	require.NoErrorf(t, driver.Quit(), "Quitting of ChromeDriver should not raise any errors.")

}
