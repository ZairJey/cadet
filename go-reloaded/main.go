package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadFile(s string) ([]string, error) {
	textBytes, err := os.ReadFile(s) //байтовый срез текста
	if err != nil {
		return nil, err
	}

	textStr := string(textBytes)          //стринговый слайс текста
	return SplitWhiteSpaces(textStr), nil //разделенный по индекс -> слово

}

func main() {

	// Отдельная функция для открытия и чтения файла
	words, err := ReadFile(os.Args[1])
	if err != nil {
		return
	}
	output := os.Args[2]

	var pars int64
	for i := 0; i < len(words); i++ {
		word := words[i]

		if len(word) != 0 && word[0] == '(' && word[len(word)-1] == ')' {
			keycommand, num := vitaskivatel(word)

			for j := i - 1; j >= 0 && num > 0; j-- {
				for _, c := range words[j] {
					if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
						num++
						break
					}
				}
				num--

				if keycommand == "cap" {
					words[j] = Capitalize(words[j])
				} else if keycommand == "up" {
					words[j] = strings.ToUpper(words[j])
				} else if keycommand == "low" {
					words[j] = strings.ToLower(words[j])
				} else if keycommand == "hex" {
					pars, err = strconv.ParseInt(words[j], 16, 32)
					if err != nil {
						fmt.Println("error")
						return
					}
					words[j] = strconv.Itoa(int(pars))
				} else if keycommand == "bin" {
					pars, err = strconv.ParseInt(words[j], 2, 32)
					if err != nil {
						fmt.Println("error")
						return
					}
					words[j] = strconv.Itoa(int(pars))
				}
			}
		}
		if word == "a" && i+1 < len(words) && len(words[i+1]) > 1 && ((words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'o' || words[i+1][0] == 'i' || words[i+1][0] == 'u' || words[i+1][0] == 'h') ||
			words[i+1][0] == 'A' || words[i+1][0] == 'E' || words[i+1][0] == 'O' || words[i+1][0] == 'I' || words[i+1][0] == 'U' || words[i+1][0] == 'H') {
			words[i] = "an"

		} else if word == "A" && i+1 < len(words) && len(words[i+1]) > 1 && ((words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'o' || words[i+1][0] == 'i' || words[i+1][0] == 'u' || words[i+1][0] == 'h') ||
			words[i+1][0] == 'A' || words[i+1][0] == 'E' || words[i+1][0] == 'O' || words[i+1][0] == 'I' || words[i+1][0] == 'U' || words[i+1][0] == 'H') {
			words[i] = "An"
		}

		for j := 0; j < len(word); j++ {
			if word[j] == '(' && word[len(word)-1] == ')' {
				words = append(words[:i], words[i+1:]...)
				i = 0
			}
		}
	}
	var localresult string
	for p := 0; p < len(words); p++ {
		if p == len(words)-1 {
			localresult += string(words[p])
		} else {
			localresult += string(words[p]) + " "
		}
	}

	var result string
	checkApostrofe := true
	for l := 0; l < len(localresult); l++ {

		if localresult[l] == ' ' {
			if checkifpunctutaion(string(localresult[l+1])) {
				result += string(localresult[l+1])
				l++
			} else {
				result += string(localresult[l])
			}

		} else {
			result += string(localresult[l])
		}

	}
	// fmt.Println(result)

	newres := ""
	for l := 0; l < len(result); l++ {

		if result[l] == '\'' {
			if checkApostrofe {
				if result[l+1] != ' ' {
					newres += string(result[l])
					checkApostrofe = false
					continue
				}
				newres += string(result[l])
				checkApostrofe = false
				l++
				continue
			} else if result[l+1] == ' ' || checkifpunctutaion(string(result[l+1])) {
				newres = newres[:len(newres)-1]
				// fmt.Println("new", newres, "s")

				newres += string(result[l])

				checkApostrofe = true
			} else {
				newres += string(result[l])
			}

		} else {
			newres += string(result[l])
		}

	}

	lan := os.WriteFile(output, []byte(newres), 0644)
	if lan != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Строка успешно записана в файл:", output)

}

func SplitWhiteSpaces(s string) []string {

	var tempStr string
	var slice []string
	flag := true

	for i := 0; i < len(s); i++ {
		// fmt.Println(string(s[i]))

		if s[i] == ' ' {
			if tempStr != "" && flag {
				slice = append(slice, tempStr)
				tempStr = ""
			}

		} else if s[i] == '\'' && (i == 0 || i-1 >= 0 && s[i-1] == ' ' || i == len(s)-1 || i+1 <= len(s)-1 && s[i+1] == ' ') {
			if tempStr != "" {
				slice = append(slice, tempStr)
				tempStr = ""
				slice = append(slice, string(s[i]))
			} else {
				slice = append(slice, string(s[i]))
			}
		} else if s[i] == '.' || s[i] == ',' || s[i] == '!' || s[i] == '?' || s[i] == ':' || s[i] == ';' {
			if tempStr != "" && !checkifpunctutaion(tempStr) {
				slice = append(slice, tempStr)
				tempStr = string(s[i])
			} else {
				tempStr += string(s[i])
			}
		} else {
			if checkifpunctutaion(tempStr) && tempStr != "" {
				slice = append(slice, tempStr)
				tempStr = ""
			}
			tempStr += string(s[i])
			if s[i] == '(' {
				flag = false
			}
			if s[i] == ')' {
				flag = true
			}
		}
	}

	if tempStr != "" {
		slice = append(slice, tempStr)
	}

	for i := 1; i < len(slice)-1; i++ {
		if slice[i] == "," && slice[i-1][0] == '(' && slice[i+1][len(slice[i+1])-1] == ')' {
			slice[i] = slice[i-1] + slice[i] + slice[i+1]
			slice = append(append(slice[0:i-1], slice[i]), slice[i+2:]...)
		}
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return (slice)
}

func checkifpunctutaion(s string) bool { //.->true
	for _, val := range s {
		if !(val == '.' || val == ',' || val == '!' || val == '?' || val == ':' || val == ';') {
			return false
		}
	}
	return true
}

func Capitalize(s string) string {
	str := []rune(s)
	isReset := true
	for ind, val := range str {
		if isReset {
			if ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9') {
				isReset = false
				continue
			} else if 'a' <= val && val <= 'z' {
				str[ind] -= 'a' - 'A'
				isReset = false
			}
		} else if !isReset && 'A' <= val && val <= 'Z' {
			str[ind] -= 'A' - 'a'
		} else if !(('a' <= val && val <= 'z') || ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9') || val == '\'') {
			isReset = true
		}
	}
	return string(str)
}

func vitaskivatel(s string) (string, int) {
	num := 1
	s2 := ""
	s1 := ""
	for _, val := range s {
		if val >= 'a' && val <= 'z' {
			s2 += string(val)
		}
		if val >= '0' && val <= '9' {
			s1 += string(val)
		}
	}
	if s1 != "" {
		s2, err := strconv.Atoi(s1)
		if err != nil {
			fmt.Println("Error")
			return "", 0
		}
		num *= s2
	}
	return s2, num
}
