class Solution {
    public int[] constructTransformedArray(int[] nums) {
        int l = nums.length;
        int[] result = new int[l];
        for (int i = 0; i < l; i++)
            result[i] = nums[((i + nums[i]) % l + l) % l];
        return result;
    }
}