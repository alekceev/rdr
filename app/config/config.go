package config

import (
	"encoding/json"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port              string `envconfig:"PORT" default:"8081" required:"true"`
	ReadTimeout       int    `envconfig:"READ_TIMEOUT" default:"30" required:"true"`
	WriteTimeout      int    `envconfig:"WRITE_TIMEOUT" default:"30" required:"true"`
	ReadHeaderTimeout int    `envconfig:"READ_HEADER_TIMEOUT" default:"30" required:"true"`
}

type VersionInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Build   string `json:"build"`
}

func (v VersionInfo) String() string {
	b, _ := json.Marshal(v)
	return string(b)
}

// init config
func Get() (Config, error) {
	config := Config{}

	if err := envconfig.Process("", &config); err != nil {
		return config, err
	}

	return config, nil
}

// Get Server host:port
func (c *Config) Addr() string {
	return ":" + c.Port
}
