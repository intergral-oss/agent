package nginx_exporter

import (
	"fmt"
	"github.com/grafana/agent/pkg/integrations"
	integrations_v2 "github.com/grafana/agent/pkg/integrations/v2"
	"github.com/grafana/agent/pkg/integrations/v2/metricsutils"
	"net/url"
)

// DefaultConfig holds the default settings for the nginx_exporter integration.
var DefaultConfig = Config{}

// Config controls the nginx_exporter integration.
type Config struct {
	ScrapeURI string `yaml:"scrape_uri"`
	NginxPlus bool   `yaml:"nginx_plus,omitempty"`
}

// UnmarshalYAML implements yaml.Unmarshaler for Config
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultConfig

	type plain Config
	return unmarshal((*plain)(c))
}

// Name returns the name of the integration that this config represents.
func (c *Config) Name() string {
	return "nginx_exporter"
}

// InstanceKey returns the address:port of the mongodb server being queried.
func (c *Config) InstanceKey(_ string) (string, error) {
	u, err := url.Parse(string(c.URI))
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}
	return u.Host, nil
}

// NewIntegration creates a new nginx_exporter
func (c *Config) NewIntegration(logger log.Logger) (integrations.Integration, error) {
	return New(logger, c)
}

func init() {
	integrations.RegisterIntegration(&Config{})
	integrations_v2.RegisterLegacy(&Config{}, integrations_v2.TypeMultiplex, metricsutils.NewNamedShim("nginx"))
}

// New creates a new nginx_exporter integration.
func New(logger log.Logger, c *Config) (integrations.Integration, error) {
	logrusLogger := NewLogger(logger)

	return integrations.NewHandlerIntegration(c.Name(), exp.Handler()), nil
}
