// Package mssql_exporter embeds https://github.com/awaragi/prometheus-mssql-exporter
package mssql_exporter //nolint:golint

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	config_util "github.com/prometheus/common/config"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grafana/agent/pkg/integrations"
	integrations_v2 "github.com/grafana/agent/pkg/integrations/v2"
	"github.com/grafana/agent/pkg/integrations/v2/metricsutils"

	"github.com/awaragi/prometheus-mssql-exporter/src"
)

// Config controls the mssql_exporter integration.
type Config struct {
	// Unique
	// DataSourceName to use to connect to mssql.
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
	return "mssql_exporter"
}

// InstanceKey returns the address:port of the mongodb server being queried.
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
	integrations_v2.RegisterLegacy(&Config{}, integrations_v2.TypeMultiplex, metricsutils.NewNamedShim("mssql"))
}

// New creates a new mssql_exporter integration. The integration scrapes
// metrics from a mssql process.
func New(log log.Logger, c *Config) (integrations.Integration, error) {
	// Unique
}