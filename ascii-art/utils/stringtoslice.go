package utils

func StringToSlice(str string) []string {
	slice := []string{}
	word := ""
	for i := 0; i < len(str); i++ {
		//if i == 0 {
		//	word = string(str[i])
		//	slice = append(slice, word)
		//	word = ""
		if str[i] == '\n' && i-1 >= 0 && str[i-1] != '\n' {
			slice = append(slice, word)
			word = ""
		} else {
			word += string(str[i])
		}
	}
	slice = append(slice, word)
	return slice
}
