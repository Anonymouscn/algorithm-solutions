class Solution {
    public int maxSideLength(int[][] mat, int threshold) {
        int result = 0;
        int rows = mat.length;
        int cols = mat[0].length;
        int[][] sum = new int[rows][cols];
        for (int r = 0, cnt = 0; r < rows; r++) {
            sum[r] = new int[cols];
            sum[r][0] = cnt + mat[r][0];
            cnt = sum[r][0];
            if (result == 0 && mat[r][0] <= threshold) {
                result = 1;
            }
        }
        for (int c = 0, cnt = 0; c < cols; c++) {
            sum[0][c] = cnt + mat[0][c];
            cnt = sum[0][c];
            if (result == 0 && mat[0][c] <= threshold) {
                result = 1;
            }
        }
        for (int r = 1; r < rows; r++) {
            for (int c = 1; c < cols; c++) {
                sum[r][c] = sum[r-1][c] + sum[r][c-1] - sum[r-1][c-1] + mat[r][c];
                if (result == 0 && mat[r][c] <= threshold) {
                    result = 1;
                }
            }
        }
        int s = rows <= cols ? rows : cols;
        for (int n = 1; n < s; n++) {
            boolean available = false;
            for (int r = n; r < rows; r++) {
                for (int c = n; c < cols; c++) {
                    int as = sum[r][c];
                    if (r > n) {
                        as -= sum[r-n-1][c];
                    }
                    if (c > n) {
                        as -= sum[r][c-n-1];
                    }
                    if (r > n && c > n) {
                        as += sum[r-n-1][c-n-1];
                    }
                    if (as <= threshold) {
                        available = true;
                        break;
                    }
                }
                if (available) {
                    break;
                }
            }
            if (available) {
                result = n + 1;
            }
        }
        return result;
    }
}