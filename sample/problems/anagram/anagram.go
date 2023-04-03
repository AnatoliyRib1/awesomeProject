package anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	records := [26]int{}
	for _, v := range s {
		records[v-'a']++
	}
	for _, v := range t {
		records[v-'a']--
	}
	for _, v := range records {
		if v != 0 {
			return false
		}
	}
	return true

}
