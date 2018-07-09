package selenium

import (
	"errors"
)

//Proxy W3C capability defines proxy configuration options.
type Proxy struct {
	ProxyType          ProxyType
	ProxyAutoconfigURL string
	FTPProxy           string
	HTTPProxy          string
	NoProxy            string
	SSLProxy           string
	SocksProxy         string
	SocksVersion       int
}

//ProxyType indicates the type of proxy configuration.
type ProxyType string

func ParseProxyType(name string) (ProxyType, error) {
	switch name {

	case string(Pac):
		return Pac, nil
	case string(Direct):
		return Direct, nil
	case string(AutoDetect):
		return AutoDetect, nil
	case string(System):
		return System, nil
	case string(Manual):
		return Manual, nil

	}

	return AutoDetect, errors.New("invalid proxy type")

}

//Constants for ProxyType
const (
	Pac        ProxyType = "pac"
	Direct     ProxyType = "direct"
	AutoDetect ProxyType = "autodetect"
	System     ProxyType = "system"
	Manual     ProxyType = "manual"
)
