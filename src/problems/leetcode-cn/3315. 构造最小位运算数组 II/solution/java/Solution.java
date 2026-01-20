class Solution {
    public int[] minBitwiseArray(List<Integer> nums) {
        return nums.stream()
                .mapToInt(s -> {
                    if ((s & 1) == 0) return -1;
                    int n = s + 1;
                    return s ^ ((n & -n) >> 1);
                })
                .toArray();
    }
}