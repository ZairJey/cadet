package ascii

import (
	"ascii-art/utils"
	"fmt"
)

func AsciiArt(slicearg []string, arraystand []string) {

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
					fmt.Print(arraystand[(val-32)*9+rune(k)])
				}
				fmt.Println()
			}
		}
	}
}
