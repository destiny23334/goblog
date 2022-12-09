package config

type Server struct {
	AppMode  string `yaml:"appMode"`
	HttpPort string `yaml:"httpPort"`
}
