package utils

import (
	"os"
	"reflect"
	"testing"

	utils "github.com/Zoooooomies/go-ls/internal"
)

func Test_GetCurrentDirectoryFiles(t *testing.T) {
	os.Chdir("../fake_directory")
	expectedResults := []string{".secret_file", "dir1", "test1.txt", "test2.txt"}

	t.Log(utils.GetCurrentDirectoryFiles())
	t.Log(expectedResults)

	if !reflect.DeepEqual(expectedResults, utils.GetCurrentDirectoryFiles()) {
		t.Fatal("Files didn't match")
	}
}
