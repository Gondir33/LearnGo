package main

import (
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type User struct {
	Name     string    `yaml:"name"`
	Age      int       `yaml:"age"`
	Comments []Comment `yaml:"comments"`
}

type Comment struct {
	Text string `yaml:"text"`
}

func writeYAML(filePath string, data []User) error {
	err := os.MkdirAll(path.Dir(filePath), 777)
	if err != nil {
		return err
	}
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bytes, 777)
}
