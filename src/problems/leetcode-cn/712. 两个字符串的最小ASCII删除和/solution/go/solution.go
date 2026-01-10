func minimumDeleteSum(s1 string, s2 string) int {
    rows, cols := len(s1), len(s2)

    dp := make([][]int, rows + 1)
    for i := 0; i <= rows; i++ {
        dp[i] = make([]int, cols + 1)
    }

    for i := 1; i <= rows; i++ {
        dp[i][0] = dp[i-1][0] + int(s1[i-1])
    }
    for i := 1; i <= cols; i++ {
        dp[0][i] = dp[0][i-1] + int(s2[i-1])
    }

    for r := 1; r <= rows; r++ {
        for c := 1; c <= cols; c++ {
            if s1[r-1] != s2[c-1] {
                dp[r][c] = min(dp[r-1][c] + int(s1[r-1]), dp[r][c-1] + int(s2[c-1]))
            } else {
                dp[r][c] = dp[r-1][c-1]
            }
        }
    }
    
    return dp[rows][cols]
}

func min(arr ...int) int {
    length := len(arr)
    if length == 0 {
        return -1
    }
    r := arr[0]
    for i := 1; i < length; i++ {
        if arr[i] < r {
            r = arr[i]
        }
    }
    return r
}