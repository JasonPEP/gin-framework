package config

import (
	"fmt"
)

type Config struct {
	Db     string `yaml:"datasource"`
	Server Server `yaml:"server"`
}
type Server struct {
	Port           uint `yaml:"port"`
	MaxConnections uint `yaml:"max-connections"`
}

var (
	Conf *Config
)

func init() {
	fmt.Print("COnfig init")
}
