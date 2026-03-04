func numSpecial(mat [][]int) int {
    rl, cl, cnt := len(mat), len(mat[0]), 0
    rc, cc := make([]int, rl), make([]int, cl)
    for i := 0; i < rl; i++ {
        for j := 0; j < cl; j++ {
            if mat[i][j] == 1 {
                rc[i]++
                cc[j]++
            }
        }
    }
    for i := 0; i < rl; i++ {
        for j := 0; j < cl; j++ {
            if mat[i][j] == 1 && rc[i] == 1 && cc[j] == 1 {
                cnt++
            }
        }
    }
    return cnt
}