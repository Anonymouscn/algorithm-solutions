class Solution {
    public int[] minBitwiseArray(List<Integer> nums) {
        int length = nums.size();
        int[] result = new int[length];
        for (int i = 0; i < length; i++) {
            result[i] = -1;
            int current = nums.get(i);
            for (int t = 0; t < current; t++) {
                if ((t | (t + 1)) == current) {
                    result[i] = t;
                    break;
                }
            }
        }
        return result;
    }
}