func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)

    for i, j := range nums {
        if v, ok := mp[j]; ok {
            return []int{v, i}
        }
        mp[target - j] = i
    }

    return []int{-1, -1}
}
