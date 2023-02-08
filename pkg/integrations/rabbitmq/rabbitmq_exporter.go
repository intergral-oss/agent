package rabbitmq

import (
	"fmt"
	"github.com/grafana/agent/pkg/integrations"
	integrations_v2 "github.com/grafana/agent/pkg/integrations/v2"
	"github.com/grafana/agent/pkg/integrations/v2/metricsutils"
	config_util "github.com/prometheus/common/config"
	"net/url"
)

// Config controls the rabbitmq_exporter integration.
type Config struct {
	RabbitURL  config_util.Secret `yaml:"rabbit_url"`
	RabbitUser config_util.Secret `yaml:"rabbit_user, omitempty"`
	RabbitPass config_util.Secret `yaml:"rabbit_pass, omitempty"`
}

// UnmarshalYAML implements yaml.Unmarshaler for Config
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Config
	return unmarshal((*plain)(c))
}

// Name returns the name of the integration that this config represents.
func (c *Config) Name() string {
	return "rabbitmq_exporter"
}

// InstanceKey returns the address:port of the mongodb server being queried.
func (c *Config) InstanceKey(_ string) (string, error) {
	u, err := url.Parse(string(c.RabbitURL))
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}
	return u.Host, nil
}

// NewIntegration creates a new rabbitmq_exporter
func (c *Config) NewIntegration(logger log.Logger) (integrations.Integration, error) {
	return New(logger, c)
}

func init() {
	integrations.RegisterIntegration(&Config{})
	integrations_v2.RegisterLegacy(&Config{}, integrations_v2.TypeMultiplex, metricsutils.NewNamedShim("rabbitmq"))
}

// New creates a new rabbitmq_exporter integration.
func New(logger log.Logger, c *Config) (integrations.Integration, error) {

}
