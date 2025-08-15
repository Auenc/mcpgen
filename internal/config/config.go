package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ProjectConfig struct {
	Title     string `yaml:"title"`
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Package   string `yaml:"package"`
	OutputDir string `yaml:"outputDir"`
}

type Config struct {
	Version string        `yaml:"version"`
	MCPGen  ProjectConfig `yaml:"mcpgen"`
	DryRun  bool          `yaml:"-"`
}

// LoadConfig reads a YAML file from the given path and unmarshals it into Config
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
