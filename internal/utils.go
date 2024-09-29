package utils

import (
	"log"
	"os"
)

func GetCurrentDirectoryFiles() []string {
	var directories []string

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		directories = append(directories, e.Name())
	}
	return directories
}
