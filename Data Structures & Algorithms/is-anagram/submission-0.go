func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	c1 := make(map[rune]int)
	c2 := make(map[rune]int)
	i2, tr := 0, []rune(t)

	for _, sv := range s {
		c1[sv] += 1
		if c1[sv] > c2[sv] {
			for i := i2; i < len(tr); i++ {
				tv := tr[i]
				c2[tv] += 1
				i2 += 1
				if sv == tv && c1[sv] <= c2[sv] {
					break
				}
			}
			if c1[sv] > c2[sv] {
				return false
			}
		}
	}

	return true
}
