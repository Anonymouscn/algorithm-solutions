impl Solution {
    pub fn min_bitwise_array(nums: Vec<i32>) -> Vec<i32> {
        return nums.iter().map(|s| {
            if (s & 1 == 0) {
                -1
            } else {
                let n = s + 1;
                s ^ ((n & -n) >> 1)
            }
        }).collect();
    }
}