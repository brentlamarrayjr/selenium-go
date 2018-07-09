package selenium

import (
	"errors"
)

//Capabilities - Selenium capabilities
type ReturnedCapabilities struct {
	BrowserName         string            `json:"browserName,omitempty"`
	BrowserVersion      string            `json:"browserVersion,omitempty"`
	PlatformName        string            `json:"platformName,omitempty"`
	AcceptInsecureCerts bool              `json:"acceptInsecureCerts,omitempty"`
	PageLoadStrategy    *PageLoadStrategy `json:"pageLoadStrategy,omitempty"`
	Proxy               *Proxy            `json:"proxy,omitempty"`
	WindowRect          bool              `json:"setWindowRect,omitempty"`
	Timeouts            *Timeouts         `json:"timeouts,omitempty"`
	Rotatable           bool              `json:"rotatable,omitempty"`

	//Chrome
	Chrome                   map[string]interface{} `json:"chrome,omitempty"`
	CSSSelectorsEnabled      bool                   `json:"cssSelectorsEnabled,omitempty"`
	DatabaseEnabled          bool                   `json:"databaseEnabled,omitempty"`
	HandlesAlerts            bool                   `json:"handlesAlerts,omitempty"`
	HasTouchScreen           bool                   `json:"hasTouchScreen,omitempty"`
	JavascriptEnabled        bool                   `json:"javascriptEnabled,omitempty"`
	LocationContextEnabled   bool                   `json:"locationContextEnabled,omitempty"`
	MobileEmulationEnabled   bool                   `json:"mobileEmulationEnabled,omitempty"`
	NativeEvents             bool                   `json:"nativeEvents,omitempty"`
	NetworkConnectionEnabled bool                   `json:"networkConnectionEnabled,omitempty"`
	Platform                 string                 `json:"platform,omitempty"`
	Version                  string                 `json:"version,omitempty"`
	TakesHeapSnapshot        bool                   `json:"takesHeapSnapshot,omitempty"`
	TakesScreenshot          bool                   `json:"takesScreenshot,omitempty"`
	UnexpectedAlertBehaviour string                 `json:"unexpectedAlertBehavior,omitempty"`
	AcceptSSLCerts           bool                   `json:"acceptSslCerts,omitempty"`
	ApplicationCacheEnabled  bool                   `json:"applicationCacheEnabled,omitempty"`
	BrowserConnectionEnabled bool                   `json:"browserConnectionEnabled,omitempty"`

	//Firefox
	AccessibilityChecks              bool   `json:"moz:accessibilityChecks,omitempty"`
	Headless                         bool   `json:"moz:headless,omitempty"`
	ProcessID                        int    `json:"moz:processID,omitempty"`
	Profile                          string `json:"moz:profile,omitempty"`
	UseNonSpecCompliantPointerOrigin bool   `json:"moz:useNonSpecCompliantPointerOrigin,omitempty"`
	WebDriverClick                   bool   `json:"moz:webdriverClick,omitempty"`
}

func NewReturnedCapabilities(capabilitiesMap map[string]interface{}) (*ReturnedCapabilities, error) {

	returnedCapabilities := &ReturnedCapabilities{}

	browserName, found := capabilitiesMap["browserName"].(string)
	if !found {
		return nil, errors.New("browserName not found")
	}
	browserVersion, found := capabilitiesMap["browserVersion"].(string)
	if !found {
		if browserName != "chrome" {
			return nil, errors.New("browserVersion not found")
		}
	}
	platformName, found := capabilitiesMap["platformName"].(string)
	if !found {
		if browserName != "chrome" {
			return nil, errors.New("platformName not found")
		}
	}
	acceptInsecureCerts, found := capabilitiesMap["acceptInsecureCerts"].(bool)
	if !found {
		return nil, errors.New("acceptInsecureCerts not found")
	}
	pageLoadStrategyString, found := capabilitiesMap["pageLoadStrategy"].(string)
	if !found {
		return nil, errors.New("pageLoadStrategy not found")
	}

	pageLoadStrategy, err := ParsePageLoadStrategy(pageLoadStrategyString)
	if err != nil {
		return nil, err
	}

	proxyMap, found := capabilitiesMap["proxy"].(map[string]interface{})
	proxy := &Proxy{}
	if found {
		proxyTypeString, found := proxyMap["proxyType"].(string)
		if !found {
			return nil, errors.New("proxy:proxyType not found")
		}

		proxyType, err := ParseProxyType(proxyTypeString)
		if err != nil {
			return nil, err
		} else {
			proxy.ProxyType = proxyType
		}

		proxyAutoconfigURL, found := proxyMap["proxyAutoconfigUrl"].(string)
		if !found {
			return nil, errors.New("proxy:proxyAutoconfigUrl not found")
		} else {
			proxy.ProxyAutoconfigURL = proxyAutoconfigURL
		}
		ftpProxy, found := proxyMap["ftpProxy"].(string)
		if !found {
			return nil, errors.New("proxy:ftpProxy not found")
		} else {
			proxy.FTPProxy = ftpProxy
		}
		httpProxy, found := proxyMap["httpProxy"].(string)
		if !found {
			return nil, errors.New("proxy:httpProxy not found")
		} else {
			proxy.HTTPProxy = httpProxy
		}
		noProxy, found := proxyMap["noProxy"].(string)
		if !found {
			return nil, errors.New("proxy:noProxy not found")
		} else {
			proxy.NoProxy = noProxy
		}
		sslProxy, found := proxyMap["sslProxy"].(string)
		if !found {
			return nil, errors.New("proxy:sslProxy not found")
		} else {
			proxy.SSLProxy = sslProxy
		}
		socksProxy, found := proxyMap["socksProxy"].(string)
		if !found {
			return nil, errors.New("proxy:socksProxy not found")
		} else {
			proxy.SocksProxy = socksProxy
		}
		socksVersion, found := proxyMap["socksVersion"].(int)
		if !found {
			return nil, errors.New("proxy:socksVersion not found")
		} else {
			proxy.SocksVersion = socksVersion
		}
	}

	timeoutsMap, found := capabilitiesMap["timeouts"].(map[string]interface{})
	timeouts := &Timeouts{}
	if found {
		script, found := timeoutsMap["script"].(float64)
		if !found {
			return nil, errors.New("timeouts:script not found")
		} else {
			timeouts.Script = int(script)
		}
		pageLoad, found := timeoutsMap["pageLoad"].(float64)
		if !found {
			return nil, errors.New("timeouts:pageLoad not found")
		} else {
			timeouts.PageLoad = int(pageLoad)
		}
		implicit, found := timeoutsMap["implicit"].(float64)
		if !found {
			return nil, errors.New("timeouts:implicit not found")
		} else {
			timeouts.Implicit = int(implicit)
		}
	} else {
		if browserName != "chrome" {
			return nil, errors.New("timeouts not found")
		}
	}

	rotatable, found := capabilitiesMap["rotatable"].(bool)
	if !found {
		return nil, errors.New("rotatable not found")
	}

	returnedCapabilities.BrowserName = browserName
	returnedCapabilities.BrowserVersion = browserVersion
	returnedCapabilities.PlatformName = platformName
	returnedCapabilities.AcceptInsecureCerts = acceptInsecureCerts
	returnedCapabilities.PageLoadStrategy = &pageLoadStrategy
	returnedCapabilities.Proxy = proxy
	returnedCapabilities.Timeouts = timeouts
	returnedCapabilities.Rotatable = rotatable

	if returnedCapabilities.BrowserName == "firefox" {

		accessibilityChecks, found := capabilitiesMap["moz:accessibilityChecks"].(bool)
		if !found {
			return nil, errors.New("accessibilityChecks not found")
		} else {
			returnedCapabilities.AccessibilityChecks = accessibilityChecks
		}

		headless, found := capabilitiesMap["moz:headless"].(bool)
		if !found {
			return nil, errors.New("headless not found")
		} else {
			returnedCapabilities.Headless = headless
		}

		processID, found := capabilitiesMap["moz:processID"].(float64)
		if !found {
			return nil, errors.New("processId not found")
		} else {
			returnedCapabilities.ProcessID = int(processID)
		}

		profile, found := capabilitiesMap["moz:profile"].(string)
		if !found {
			return nil, errors.New("profile not found")
		} else {
			returnedCapabilities.Profile = profile
		}

		useNonSpecCompliantPointerOrigin, found := capabilitiesMap["moz:useNonSpecCompliantPointerOrigin"].(bool)
		if !found {
			return nil, errors.New("useNonSpecCompliantPointerOrigin not found")
		} else {
			returnedCapabilities.UseNonSpecCompliantPointerOrigin = useNonSpecCompliantPointerOrigin
		}

		webDriverClick, found := capabilitiesMap["moz:webdriverClick"].(bool)
		if !found {
			return nil, errors.New("webdriverClick not found")
		} else {
			returnedCapabilities.WebDriverClick = webDriverClick
		}

	}

	if returnedCapabilities.BrowserName == "chrome" {

		chrome, found := capabilitiesMap["chrome"].(map[string]interface{})
		if !found {
			return nil, errors.New("chrome not found")
		} else {
			returnedCapabilities.Chrome = chrome
		}

		windowRect, found := capabilitiesMap["setWindowRect"].(bool)
		if !found {
			return nil, errors.New("setWindowRect not found")
		} else {
			returnedCapabilities.WindowRect = windowRect
		}

		cssSelectorsEnabled, found := capabilitiesMap["cssSelectorsEnabled"].(bool)
		if !found {
			return nil, errors.New("cssSelectorsEnabled not found")
		} else {
			returnedCapabilities.CSSSelectorsEnabled = cssSelectorsEnabled
		}

		databaseEnabled, found := capabilitiesMap["databaseEnabled"].(bool)
		if !found {
			return nil, errors.New("databaseEnabled not found")
		} else {
			returnedCapabilities.DatabaseEnabled = databaseEnabled
		}

		handlesAlerts, found := capabilitiesMap["handlesAlerts"].(bool)
		if !found {
			return nil, errors.New("handlesAlerts not found")
		} else {
			returnedCapabilities.HandlesAlerts = handlesAlerts
		}

		hasTouchScreen, found := capabilitiesMap["hasTouchScreen"].(bool)
		if !found {
			return nil, errors.New("hasTouchScreen not found")
		} else {
			returnedCapabilities.HasTouchScreen = hasTouchScreen
		}

		locationContextEnabled, found := capabilitiesMap["locationContextEnabled"].(bool)
		if !found {
			return nil, errors.New("locationContextEnabled not found")
		} else {
			returnedCapabilities.LocationContextEnabled = locationContextEnabled
		}

		mobileEmulationEnabled, found := capabilitiesMap["mobileEmulationEnabled"].(bool)
		if !found {
			return nil, errors.New("mobileEmulationEnabled not found")
		} else {
			returnedCapabilities.MobileEmulationEnabled = mobileEmulationEnabled
		}

		nativeEvents, found := capabilitiesMap["nativeEvents"].(bool)
		if !found {
			return nil, errors.New("nativeEvents not found")
		} else {
			returnedCapabilities.NativeEvents = nativeEvents
		}

		networkConnectionEnabled, found := capabilitiesMap["networkConnectionEnabled"].(bool)
		if !found {
			return nil, errors.New("networkConnectionEnabled not found")
		} else {
			returnedCapabilities.NetworkConnectionEnabled = networkConnectionEnabled
		}

		platform, found := capabilitiesMap["platform"].(string)
		if !found {
			return nil, errors.New("platform not found")
		} else {
			returnedCapabilities.Platform = platform
		}

		version, found := capabilitiesMap["version"].(string)
		if !found {
			return nil, errors.New("version not found")
		} else {
			returnedCapabilities.Version = version
		}

		takesScreenshot, found := capabilitiesMap["takesScreenshot"].(bool)
		if !found {
			return nil, errors.New("takesScreenshot not found")
		} else {
			returnedCapabilities.TakesScreenshot = takesScreenshot
		}

		takesHeapSnapshot, found := capabilitiesMap["takesHeapSnapshot"].(bool)
		if !found {
			return nil, errors.New("takesHeapSnapshot not found")
		} else {
			returnedCapabilities.TakesHeapSnapshot = takesHeapSnapshot
		}

		unexpectedAlertBehaviour, found := capabilitiesMap["unexpectedAlertBehaviour"].(string)
		if !found {
			return nil, errors.New("unexpectedAlertBehaviour not found")
		} else {
			returnedCapabilities.UnexpectedAlertBehaviour = unexpectedAlertBehaviour
		}

		acceptSSLCerts, found := capabilitiesMap["acceptSslCerts"].(bool)
		if !found {
			return nil, errors.New("acceptSslCerts not found")
		} else {
			returnedCapabilities.AcceptSSLCerts = acceptSSLCerts
		}

		applicationCacheEnabled, found := capabilitiesMap["applicationCacheEnabled"].(bool)
		if !found {
			return nil, errors.New("applicationCacheEnabled not found")
		} else {
			returnedCapabilities.ApplicationCacheEnabled = applicationCacheEnabled
		}

		browserConnectionEnabled, found := capabilitiesMap["browserConnectionEnabled"].(bool)
		if !found {
			return nil, errors.New("browserConnectionEnabled not found")
		} else {
			returnedCapabilities.BrowserConnectionEnabled = browserConnectionEnabled
		}

	}

	return returnedCapabilities, nil

}
