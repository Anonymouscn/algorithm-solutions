impl Solution {
    pub fn is_trionic(nums: Vec<i32>) -> bool {
        if nums.len() < 4 {
            return false;
        }

        fn climb(nums: &[i32], mut i: usize, cmp: fn(i32, i32) -> bool) -> Option<usize> {
            let start = i;
            while i + 1 < nums.len() && cmp(nums[i + 1], nums[i]) {
                i += 1;
            }
            if i > start { Some(i) } else { None }
        }

        let mut i = 0;

        i = match climb(&nums, i, |a, b| a > b) { Some(x) => x, None => return false };
        i = match climb(&nums, i, |a, b| a < b) { Some(x) => x, None => return false };
        i = match climb(&nums, i, |a, b| a > b) { Some(x) => x, None => return false };

        i == nums.len() - 1
    }
}