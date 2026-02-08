func minPairSum(nums []int) int {
    length, result := len(nums), 0
    sort.Ints(nums)
    for i := 0; i < length / 2; i++ {
        if sum := nums[i] + nums[length-1-i]; sum > result {
            result = sum
        }
    }
    return result
}