package config

type Database struct {
	Db     string `yaml:"db"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Name   string `yaml:"name"`
}
