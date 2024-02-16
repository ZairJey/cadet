package utils

import (
	"os"
	"strings"
)

func ReadFile() ([]string, error) {
	standard, err := os.ReadFile("standard.txt")
	if err != nil {
		return nil, err
	}
	chars := strings.Split(string(standard), "\n")
	return chars, err
}
