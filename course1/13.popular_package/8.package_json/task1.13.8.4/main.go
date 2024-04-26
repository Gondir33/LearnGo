package main

import (
	"encoding/json"
	"os"
	"path"
)

func writeJSON(filePath string, data interface{}) error {

	err := os.MkdirAll(path.Dir(filePath), 777)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bytes, 777)
}

func main() {
	// Write JSON data to file
	data := []map[string]interface{}{
		{
			"name": "Elliot",
			"age":  25,
		},
		{
			"name": "Fraser",
			"age":  30,
		},
	}
	err := writeJSON("users.json", data)
	if err != nil {
		panic(err)
	}
}
