class Solution {

    List<List<Integer>> result = new LinkedList<>();
    List<Integer> path = new LinkedList<>();

    public List<List<Integer>> subsetsWithDup(int[] nums) {
        Arrays.sort(nums);
        backtrack(nums, 0);
        return result;
    }

    public void backtrack(int[] nums, int i) {
        if (i == nums.length) {
            result.add(new LinkedList<>(path));
            return;
        }

        int j = i;
        while (j < nums.length && nums[j] == nums[i]) j++;
        int cnt = j - i;

        backtrack(nums, j);

        for (int t = 0; t < cnt; t++) {
            path.add(nums[i]);
            backtrack(nums, j);
        }

        path = path.subList(0, path.size()-cnt);
    }
}