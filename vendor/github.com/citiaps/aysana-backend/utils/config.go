package utils

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
)

var (
	Config Configuration
)

type MongoDB struct {
	Host     string `yaml:"host"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Pass     string `yaml:"pass"`
}

type Elastic struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Pass     string `yaml:"pass"`
	Name     string `yaml:"name"`
	Index    string `yaml:"index"`
	Type     string `yaml:"type"`
}

type Server struct {
	Port string `yaml:port`
}

type Configuration struct {
	MongoDB MongoDB `yaml:"mongodb"`
	Elastic Elastic `yaml:"elastic"`
	Server  Server  `yaml:"server"`
}

func LoadConfig(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	Check(err)

	err = yaml.Unmarshal(bytes, &Config)
	Check(err)
}
