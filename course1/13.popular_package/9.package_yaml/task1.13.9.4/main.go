package main

import (
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

func writeYAML(filePath string, data interface{}) error {
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
