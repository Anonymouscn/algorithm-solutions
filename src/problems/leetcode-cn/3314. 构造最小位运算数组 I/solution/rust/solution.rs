impl Solution {
    pub fn min_bitwise_array(nums: Vec<i32>) -> Vec<i32> {
        let mut result : Vec<i32> = Vec::new();
        for item in nums {
            result.push(-1);
            for t in 0..item {
                if ((t | (t + 1)) == item) {
                    let mut l = result.len();
                    result[l-1] = t;
                    break;
                }
            }
        }
        return result;
    }
}