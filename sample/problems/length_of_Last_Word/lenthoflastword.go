package length_of_Last_Word

import "strings"

func lengthOfLastWord(s string) int {
	sum := 0
	s = strings.TrimSpace(s)
	for i := len(s) - 1; i >= +0; i-- {
		if s[i] != ' ' {
			sum++
		} else {
			break
		}

	}
	return sum
}
