class Solution {
    public void moveZeroes(int[] nums) {
        int length = nums.length;
        if (length < 2) return;
        int slow = 0, fast = 0;
        for (; fast < length; fast++) {
            if (nums[fast] != 0) {
                if (fast != slow) {
                    nums[slow] ^= nums[fast];
                    nums[fast] ^= nums[slow];
                    nums[slow] ^= nums[fast];
                }
                slow++;
            }
        }
    }
}