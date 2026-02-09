impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let length = nums.len();
        if length < 2 {
            return;
        }

        let (mut slow, mut fast) = (0, 0);
        while fast < length {
            if nums[fast] != 0 {
                if fast != slow {
                    (nums[slow], nums[fast]) = (nums[fast], nums[slow]);
                }
                slow += 1;
            }
            fast += 1;
        }
    }
}