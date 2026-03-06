func twoSum(nums []int, target int) []int {
    length, m := len(nums), make(map[int]int)
    for i := 0; i < length; i++ {
        t := target - nums[i]
        if v, ok := m[t]; ok {
            return []int{v, i}
        } else {
            m[nums[i]] = i
        }
    }
    return []int{0, 0}
}