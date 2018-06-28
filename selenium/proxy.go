package selenium

//Proxy W3C capability defines proxy configuration options.
type Proxy struct {
	proxyType          ProxyType
	proxyAutoconfigURL string
	ftpProxy           string
	httpProxy          string
	noProxy            string
	sslProxy           string
	socksProxy         string
	socksVersion       int
}

//ProxyType indicates the type of proxy configuration.
type ProxyType string

//Constants for ProxyType
const (
	Pac        ProxyType = "pac"
	Direct     ProxyType = "direct"
	AutoDetect ProxyType = "autodetect"
	System     ProxyType = "system"
	Manual     ProxyType = "manual"
)
