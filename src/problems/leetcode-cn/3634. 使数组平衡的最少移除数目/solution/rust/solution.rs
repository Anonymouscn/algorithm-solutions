impl Solution {
    pub fn min_removal(mut nums: Vec<i32>, k: i32) -> i32 {
        let length = nums.len();
        if length < 2 {
            return 0;
        }

        nums.sort_unstable();
        let (mut l, mut r, mut inc) = (0, 1, 0);

        while r < length {
            if nums[r] as i64 <= (nums[l] as i64 * k as i64) {
                inc = inc.max(r-l+1);
                r += 1;
            } else {
                l += 1;
            }
        }

        (length-inc) as i32
    }
}