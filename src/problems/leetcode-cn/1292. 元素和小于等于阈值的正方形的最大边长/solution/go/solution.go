func maxSideLength(mat [][]int, threshold int) int {
    result, rows, cols := 0, len(mat), len(mat[0])
    sum := make([][]int, rows)
    // 1.计算二维前缀和
    for r, cnt := 0, 0; r < rows; r++ {
        sum[r] = make([]int, cols)
        sum[r][0] = cnt + mat[r][0]
        cnt = sum[r][0]
        if result == 0 && mat[r][0] <= threshold {
            result = 1
        }
    }
    for c, cnt := 0, 0; c < cols; c++ {
        sum[0][c] = cnt + mat[0][c]
        cnt = sum[0][c]
        if result == 0 && mat[0][c] <= threshold {
            result = 1
        }
    }
    for r := 1; r < rows; r++ {
        for c := 1; c < cols; c++ {
            sum[r][c] = sum[r][c-1] + sum[r-1][c] - sum[r-1][c-1] + mat[r][c]
            if result == 0 && mat[r][c] <= threshold {
                result = 1
            }
        }
    }
    // 2.正方形边长定义域: [1, min(m, n)]
    for n := 1; n < rows; n++ {
        available := false
        for r := n; r < rows; r++ {
            for c := n; c < cols; c++ {
                s := sum[r][c]
                if r - n > 0 {
                    s -= sum[r-n-1][c]
                }
                if c - n > 0 {
                    s -= sum[r][c-n-1]
                }
                if r -n > 0 && c - n > 0 {
                    s += sum[r-n-1][c-n-1]
                }
                if s <= threshold {
                    available = true
                    result = n + 1
                    break
                }
            }
            if available {
                break
            }
        }
    }
    return result
}