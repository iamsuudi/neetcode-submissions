func hasDuplicate(nums []int) bool {
    mp := make(map[int]bool)
    for _, n := range nums {
        if v, check := mp[n]; check && v {
            return v
        } else {
            mp[n] = true
        }
    }
    return false
}
