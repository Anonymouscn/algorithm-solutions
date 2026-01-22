func minimumTotal(triangle [][]int) int {
    rows := len(triangle)
    for r := rows - 2; r >= 0; r-- {
        cols := len(triangle[r])
        for c := 0; c < cols; c++ {
            triangle[r][c] += min(triangle[r+1][c], triangle[r+1][c+1])
        }
    }
    return triangle[0][0]
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}