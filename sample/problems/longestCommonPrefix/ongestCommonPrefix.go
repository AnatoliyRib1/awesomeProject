package longestCommonPrefix

func longestCommonPrefix(strs []string) string {

	for i, j := range strs[0] {
		for _, k := range strs[1:] {
			if i == len(k) || byte(j) != k[i] {
				return strs[0][:i]
			}
		}

	}
	return strs[0]
}
