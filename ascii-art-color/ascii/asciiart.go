package ascii

import (
	"ascii-art/utils"
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Orange = "\033[38;5;208m"
	Black  = "\033[30m"
)

func AsciiArt(slicearg []string, arraystand []string, arg1 string, str string) {

	color := ""
	if len(arg1) > 8 {
		if str != "" {
			if arg1[0:8] == "--color=" {
				color = arg1[8:]
				if Colorize1(color) == "notfound" {
					fmt.Println("Wrong color")
					return
				}
			} else {
				fmt.Println("Wrong format, use for EX: go run . --color=<color> <letters to be colored> \"something\"")
				return
			}
		}
	}

	for i, text := range slicearg {
		if text == "" {
			if !utils.CheckisEmpty(slicearg) {
				fmt.Println()
			} else if i != 0 && utils.CheckisEmpty(slicearg) {
				fmt.Println()
			}
		} else {
			for k := 1; k < 9; k++ {
				for _, val := range text {
					if SeemsLike(string(val), str) {
						fmt.Print(Colorize(arraystand[(val-32)*9+rune(k)], color))
					} else {
						fmt.Print(arraystand[(val-32)*9+rune(k)])
					}
				}
				fmt.Println()
			}
		}
	}
}
func SeemsLike(slice string, str string) bool {
	for _, val := range slice {
		for _, char := range str {
			if val == char {
				return true
			}
		}
	}
	return false
}

func Colorize(text string, colorName string) string {
	var colorCode string

	// Установка кода цвета в зависимости от предоставленного имени цвета
	switch colorName {
	case "red":
		colorCode = Red
	case "green":
		colorCode = Green
	case "blue":
		colorCode = Blue
	case "yellow":
		colorCode = Yellow
	case "cyan":
		colorCode = Cyan
	case "purple":
		colorCode = Purple
	case "white":
		colorCode = White
	case "orange":
		colorCode = Orange
	case "black":
		colorCode = Black

	default:
		//fmt.Println("Неподдерживаемый цвет. Установлен белый цвет по умолчанию.")
		colorCode = "notfound"
	}

	return colorCode + text + Reset
}

func Colorize1(text string) string {
	var colorCode string

	// Установка кода цвета в зависимости от предоставленного имени цвета
	switch text {
	case "red":
		colorCode = Red
	case "green":
		colorCode = Green
	case "blue":
		colorCode = Blue
	case "yellow":
		colorCode = Yellow
	case "cyan":
		colorCode = Cyan
	case "purple":
		colorCode = Purple
	case "white":
		colorCode = White
	case "orange":
		colorCode = Orange
	case "black":
		colorCode = Black

	default:
		//fmt.Println("Неподдерживаемый цвет. Установлен белый цвет по умолчанию.")
		colorCode = "notfound"
	}

	return colorCode
}
