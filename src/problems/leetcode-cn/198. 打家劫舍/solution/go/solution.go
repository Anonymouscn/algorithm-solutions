func rob(nums []int) int {
    length := len(nums)
    if length == 1 {
        return nums[0]
    }

    dp := make([]int, length)
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])

    for i := 2; i < length; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
    }

    return dp[length-1]
}