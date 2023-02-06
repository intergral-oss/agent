// Package template_exporter embeds https://github.com/prometheus/template_exporter
package observability_agent //nolint:golint

import (
	"context"
	"fmt"
	"os"
	"strings"
	"net/url"

	config_util "github.com/prometheus/common/config"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grafana/agent/pkg/integrations"
	integrations_v2 "github.com/grafana/agent/pkg/integrations/v2"
	"github.com/grafana/agent/pkg/integrations/v2/metricsutils"

	"github.com/prometheus/template_exporter"
)

// Config controls the template_exporter integration.
type Config struct {
	// Unique
	// DataSourceName to use to connect to template server.
	DataSourceName config_util.Secret `yaml:"data_source_name,omitempty"`

	//Add any other config file options here

}

// UnmarshalYAML implements yaml.Unmarshaler for Config.
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Config
	return unmarshal((*plain)(c))
}

// Name returns the name of the integration this config is for.
func (c *Config) Name() string {
	return "template_exporter"
}

// InstanceKey returns the address:port of the template server being queried.
func (c *Config) InstanceKey(_ string) (string, error) {
	// Unique
	// Use this to retrieve the datasource name from the config
}

// NewIntegration converts this config into an instance of a configuration.
func (c *Config) NewIntegration(l log.Logger) (integrations.Integration, error) {
	return New(l, c)
}

func init() {
	integrations.RegisterIntegration(&Config{})
	integrations_v2.RegisterLegacy(&Config{}, integrations_v2.TypeMultiplex, metricsutils.NewNamedShim("template"))
}

// New creates a new template_exporter integration. The integration scrapes
// metrics from a template process.
func New(log log.Logger, c *Config) (integrations.Integration, error) {
	// Unique
	// Use to implement go methods from the template exporter
	// This starts an instance of that exporter
}
