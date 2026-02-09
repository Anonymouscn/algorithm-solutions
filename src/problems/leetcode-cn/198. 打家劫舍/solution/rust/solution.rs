impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let length = nums.len();
        if length == 1 {
            return nums[0];
        }

        let mut dp = vec![0i32;length];
        dp[0] = nums[0];
        dp[1] = nums[0].max(nums[1]);

        for i in 2..length {
            dp[i] = dp[i-1].max(dp[i-2]+nums[i])
        }

        return dp[length-1]
    }
}