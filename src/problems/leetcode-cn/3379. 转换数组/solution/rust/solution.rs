impl Solution {
    pub fn construct_transformed_array(nums: Vec<i32>) -> Vec<i32> {
        let (n, mut result) = (nums.len() as i32, vec![0; nums.len()]);
        for i in 0..nums.len() {
            result[i] = nums[(((((i as i32)+nums[i])%n)+n)%n) as usize];
        }
        result
    }
}