class Solution {
    int start = 0;
    Integer[] arr;

    public boolean isTrionic(int[] nums) {
        if (nums.length < 4) return false;
        arr = Arrays.stream(nums).boxed().toArray(Integer[]::new);
        if (!climb((i1, i2) -> i1-i2) || !climb((i1, i2) -> i2-i1) || !climb((i1, i2) -> i1-i2)) return false;
        return start == nums.length-1;
    }

    boolean climb(Comparator<Integer> comparator) {
        int pre = start;
        for (; start + 1 < arr.length && comparator.compare(arr[start+1], arr[start]) > 0; start++) {}
        return start > pre;
    }
}