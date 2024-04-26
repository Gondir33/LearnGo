package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, prefix string, isLast bool, depth int) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		childPath := filepath.Join(path, file.Name())

		var newPrefix, currentPrefix string

		if i == len(files)-1 {
			currentPrefix = prefix + "└─── "
			newPrefix = prefix + "	"
		} else {
			currentPrefix = prefix + "├─── "
			newPrefix = prefix + "│	"
		}

		fmt.Print(currentPrefix + file.Name())

		if file.IsDir() {
			if depth > 0 {
				fmt.Println()
				printTree(childPath, newPrefix, i == len(files)-1, depth-1)
			} else {
				fmt.Println("...")
			}
		} else {
			fmt.Println()
		}
	}
}

func main() {
	// получение флага
	var depth int
	flag.IntVar(&depth, "n", 0, "Целочисленное значение")
	flag.Parse()
	path := flag.Args()[0]
	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}

	printTree(path, "", false, depth)
}
