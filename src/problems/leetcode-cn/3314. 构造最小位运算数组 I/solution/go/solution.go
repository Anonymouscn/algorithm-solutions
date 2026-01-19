func minBitwiseArray(nums []int) []int {
    length := len(nums)
    result := make([]int, length)
    for i := 0; i < length; i ++ {
        result[i] = -1
        for t := 0; t < nums[i]; t++ {
            if (t | (t + 1) == nums[i]) {
                result[i] = t
                break
            }
        }
    }
    return result
}