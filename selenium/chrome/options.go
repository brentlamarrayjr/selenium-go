package chrome

type ChromeOptions struct {
	Args             []string               `json:",omitempty"`
	Binary           string                 `json:",omitempty"`
	Extensions       []string               `json:",omitempty"`
	LocalState       map[string]interface{} `json:",omitempty"`
	Prefs            map[string]interface{} `json:",omitempty"`
	Detach           bool                   `json:",omitempty"`
	DebuggerAddress  string                 `json:",omitempty"`
	ExcludeSwitches  []string               `json:",omitempty"`
	MinidumpPath     string                 `json:",omitempty"`
	MobileEmulation  map[string]interface{} `json:",omitempty"`
	PerfLoggingPrefs *PerfLoggingPrefs      `json:",omitempty"`
	windowTypes      []string               `json:",omitempty"`
}

func (options *ChromeOptions) AddArgs(args ...string) {
	options.Args = append(options.Args, args...)
}

func (options *ChromeOptions) SetBinary(binary string) {
	options.Binary = binary
}

func (options *ChromeOptions) AddExtensions(extensions ...string) {
	options.Extensions = append(options.Extensions, extensions...)
}

func (options *ChromeOptions) AddLocalState(name string, value interface{}) {
	options.LocalState[name] = value
}

func (options *ChromeOptions) AddPref(name string, value interface{}) {
	options.Prefs[name] = value
}

func (options *ChromeOptions) SetDetach(detach bool) {
	options.Detach = detach
}

func (options *ChromeOptions) SetDebuggerAddress(address string) {
	options.DebuggerAddress = address
}

func (options *ChromeOptions) AddExcludeSwitches(switches ...string) {
	options.ExcludeSwitches = append(options.ExcludeSwitches, switches...)
}

func (options *ChromeOptions) SetMinidumpPath(path string) {
	options.MinidumpPath = path
}

func (options *ChromeOptions) AddMobileEmulation(name string, value interface{}) {
	options.MobileEmulation[name] = value
}

func (options *ChromeOptions) SetPerfLoggingPrefs(prefs *PerfLoggingPrefs) {
	options.PerfLoggingPrefs = prefs
}

func (options *ChromeOptions) AddWindowTypes(types ...string) {
	options.windowTypes = append(options.windowTypes, types...)
}

//PerfLoggingPrefs specifies performance logging preferences
type PerfLoggingPrefs struct {
	EnableNetwork                bool   `json:",omitempty"`
	EnablePage                   bool   `json:",omitempty"`
	TraceCategories              string `json:",omitempty"`
	BufferUsageReportingInterval int    `json:",omitempty"`
}
