package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile() ([]string, int) {

	standard, err := os.ReadFile("standard.txt")
	if err != nil {
		return nil, 1
	}
	if len(standard) != 6623 {
		fmt.Println("wrong size standard")
		return nil, 1
	}
	chars := strings.Split(string(standard), "\n")
	return chars, 2
}
