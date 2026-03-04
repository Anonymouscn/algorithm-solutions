func findMaxConsecutiveOnes(nums []int) int {
    result, length, l, r := 0, len(nums), 0, 0
    for l < length && r < length {
        if nums[l] == 1 {
            for r = l; r < length-1 && nums[r+1] == 1; r++ {}
            result = max(result, r-l+1)
            l = r + 1
        } else {
            l++
        }
    }
    return result
}