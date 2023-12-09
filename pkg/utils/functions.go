package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func SetRepoName() {
	s, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory")
		return
	}
	REPONAME = filepath.Base(s)
}

// Check if a string array contains a string
func ArrContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
