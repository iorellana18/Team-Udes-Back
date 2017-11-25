package utils

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
)

var (
	Config Configuration
)

type Postgres struct {
	Rdbms   string `yaml:"rdbms"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	Ip      string `yaml:"ip"`
	Port    string `yaml:"port"`
	Name    string `yaml:"name"`
	Sslmode string `yaml:"sslmode"`
}

type Elastic struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Index    string `yaml:"index"`
	Type     string `yaml:"type"`
}

type Server struct {
	Port string `yaml:port`
}

type Configuration struct {
	Postgres Postgres `yaml:"postgres"`
	Elastic Elastic `yaml:"elastic"`
	Server  Server  `yaml:"server"`
}

func LoadConfig(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	Check(err)

	err = yaml.Unmarshal(bytes, &Config)
	Check(err)
}
