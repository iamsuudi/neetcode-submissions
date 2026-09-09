import "slices"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		slices.Sort(runes)
		sorted := string(runes)
		mp[sorted] = append(mp[sorted], s)
	}

	result := make([][]string, 0, len(mp))

	for _, v := range mp {
		result = append(result, v)
	}

	return result
}
