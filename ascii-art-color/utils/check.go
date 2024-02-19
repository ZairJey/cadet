package utils

func CheckisEmpty(slicearg []string) bool {
	for i := range slicearg {
		if slicearg[i] != "" {
			return false
		}
	}
	return true
}
