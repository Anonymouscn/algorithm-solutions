class Solution {
    public int minimumDeleteSum(String s1, String s2) {
        int rows = s1.length(), cols = s2.length();

        int[][] dp = new int[rows+1][];
        for (int i = 0; i <= rows; i++) {
            dp[i] = new int[cols+1];
        }

        for (int i = 1; i <= rows; i++) {
            dp[i][0] = dp[i-1][0] + s1.charAt(i-1);
        }
        for (int i = 1; i <= cols; i++) {
            dp[0][i] = dp[0][i-1] + s2.charAt(i-1);
        }

        for (int r = 1; r <= rows; r++) {
            for (int c = 1; c <= cols; c++) {
                if (s1.charAt(r-1) != s2.charAt(c-1)) {
                    int l = dp[r-1][c] + s1.charAt(r-1), u = dp[r][c-1] + s2.charAt(c-1);
                    dp[r][c] = l < u ? l : u;
                } else {
                    dp[r][c] = dp[r-1][c-1];
                }
            }
        }

        return dp[rows][cols];
    }
}