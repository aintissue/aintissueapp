package app

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Config stores configuration vars
type Config struct {
	Dev   bool   `yaml:"dev"`
	DbURI string `yaml:"db_uri"`
}

// Load method loads configuration file to Config struct
func (c *Config) load(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		loge(err)
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		loge(err)
	}
}

// Initializes configuration
func initConfig() *Config {
	c := &Config{}
	c.load("data/config.yaml")
	return c
}
