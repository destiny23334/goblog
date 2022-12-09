package config

type Config struct {
	Server   `yaml:"server"`
	Database `yaml:"database"`
}

var GlobalConfig Config
