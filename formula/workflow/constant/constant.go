package constant

const (
	BootTimeoutDefault   = 30000
	ClientMonitorDefault = 15000
	ConfigFileName       = "configuration.toml"
	ApiTriggerRoute      = "/api/v1/trigger"
	ApiPingRoute         = "/api/v1/ping"
	ApiMetricsRoute      = "/api/v1/metrics"
	ApiVersionRoute      = "/api/version"
	LogDurationKey       = "duration"
)

const (
	ContentType     = "Content-Type"
	ContentTypeCBOR = "application/cbor"
	ContentTypeJSON = "application/json"
	ContentTypeYAML = "application/x-yaml"
)

// ApplicationVersion indicates the version of the application itself, not the SDK - will be overwritten by build
var ApplicationVersion string = "0.0.0"
