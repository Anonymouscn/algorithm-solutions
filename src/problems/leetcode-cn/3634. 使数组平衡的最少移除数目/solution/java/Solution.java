class Solution {
    public int minRemoval(int[] nums, int k) {
        int length = nums.length;
        if (length < 2) return 0;

        Arrays.sort(nums);

        int l = 0, r = 1, inc = 0;
        while (r < length) {
            if ((long) nums[r] <= (long) nums[l]* (long) k) {
                int t = r-l+1;
                inc = t > inc ? t : inc;
                r++;
            } else {
                l++;
            }
        }

        return length-inc;
    }
}