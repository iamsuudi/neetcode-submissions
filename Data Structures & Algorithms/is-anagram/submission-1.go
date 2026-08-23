func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	count := make(map[rune]int)
	p, tr := 0, []rune(t)

	for _, sv := range s {
		count[sv]++
		if count[sv] > 0 {
			for i := p; i < len(tr); i++ {
				tv := tr[i]
				count[tv]--
				p++
				if count[sv] <= 0 {
					break
				}
			}
			if count[sv] > 0 {
				return false
			}
		}
	}

	return true
}
