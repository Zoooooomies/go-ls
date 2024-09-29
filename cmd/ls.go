package main

import (
	"fmt"

	utils "github.com/Zoooooomies/go-ls/internal"
)

func main() {
	directories := utils.GetCurrentDirectoryFiles()
	fmt.Println(directories)
}
