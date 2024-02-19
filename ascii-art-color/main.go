package main

import (
	"ascii-art/ascii"
	"ascii-art/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	standart, err := utils.ReadFile()
	if err != 2 {
		fmt.Println("error")
		return
	}

	color := "--color="
	if len(os.Args) == 3 || len(os.Args) == 4 {

		for i := 0; i < len(color); i++ {
			if len(os.Args[1]) > len(color) {
				if color[i] == os.Args[1][i] {
					continue
				}
			} else {
				fmt.Println("Wrong format, use for EX: go run . --color=<color> <letters to be colored> \"something\"")
				return
			}
		}
	}

	if len(os.Args) > 4 || len(os.Args) == 1 {
		return
	}

	if len(os.Args) > 2 {
		for _, val := range os.Args[len(os.Args)-1] {
			if val != 10 && val < 32 || val > 126 {
				fmt.Println("wrong symbols in argument")
				return
			}
		}

		arg := strings.Replace(os.Args[len(os.Args)-1], "\\n", "\n", -1)
		slicearg := strings.Split(arg, "\n")
		if len(os.Args) == 3 || len(os.Args) == 4 {
			ascii.AsciiArt(slicearg, standart, os.Args[1], os.Args[2])
		}
	} else {
		char := ""
		for _, val := range os.Args[1] {
			if val != 10 && val < 32 || val > 126 {
				fmt.Println("wrong symbols in argument")
				return
			}
		}

		arg := strings.Replace(os.Args[1], "\\n", "\n", -1)
		slicearg := strings.Split(arg, "\n")

		ascii.AsciiArt(slicearg, standart, os.Args[1], char)
	}
}
