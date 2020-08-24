package freq_alphabets

import "strconv"

func freqAlphabets(s string) string {
	result := ""
	lenS := len(s)
	for idx := lenS - 1; idx >= 0; {
		if s[idx] == '#' {
			temp := string(s[idx-2]) + string(s[idx-1])
			tempResult, _ := strconv.Atoi(temp)
			result = string('a'+tempResult-1) + result
			idx -= 3
		} else {
			temp := string(s[idx])
			tempResult, _ := strconv.Atoi(temp)
			result = string('a'+tempResult-1) + result
			idx -= 1
		}
	}
	return result
}
