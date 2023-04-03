package canConstruct

func canConstruct(ransomNote string, magazine string) bool {

	res1 := [26]int{}
	res2 := [26]int{}
	for _, v := range ransomNote {
		res1[v-'a']++
	}
	for _, v := range magazine {
		res2[v-'a']++
	}
	for i := 0; i < len(res1); i++ {
		if res1[i] > res2[i] {
			return false
		}
	}
	return true

}
