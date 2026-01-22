class Solution {
    List<String> result = new LinkedList<>();

    public List<String> generateParenthesis(int n) {
        backtrack("", 0, 0, n);
        return result;
    }

    public void backtrack(String path, int open, int close, int max) {
        if (path.length() == 2*max) {
            result.add(path);
            return;
        }

        if (open < max) {
            backtrack(path+"(", open+1, close, max);
        }
        if (close < open) {
            backtrack(path+")", open, close+1, max);
        }
    }
}