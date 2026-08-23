func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	
	count := make(map[rune]int)

	for _, v := range s {
		count[v]++
	}

	for _, v := range t {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}

	return true
}
