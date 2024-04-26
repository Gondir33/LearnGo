package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getYAML(data []Config) (string, error) {
	res, err := yaml.Marshal(data)
	return string(res), err
}

func main() {
	data := []Config{{
		Server: Server{"8080"},
		Db: Db{
			Host:     "localhost",
			Port:     "5432",
			User:     "admin",
			Password: "password123",
		}},
	}
	s, err := getYAML(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
