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
	if err != nil {
		fmt.Errorf("error")
		return
	}
	if len(os.Args) != 2 {
		fmt.Println("You can write only 1 argument")
		return
	}

	arg := strings.Replace(os.Args[1], "\\n", "\n", -1)
	slicearg := strings.Split(arg, "\n")
	for _, val := range arg {
		if val != 10 && val < 32 || val > 126 {
			fmt.Println("wrong symbols in argument")
			return
		}
	}

	ascii.AsciiArt(slicearg, standart)
}
