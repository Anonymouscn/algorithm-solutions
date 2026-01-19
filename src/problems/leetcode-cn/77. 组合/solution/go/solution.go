func combine(n int, k int) [][]int {
    result, selected := make([][]int, 0), make([]int, 0)
    var backtrack func(start, k int)
    backtrack = func(start, k int) {
        if len(selected) == k {
            tmp := make([]int, len(selected))
            copy(tmp, selected)
            result = append(result, tmp)
            return
        }
        for i := start; i <= n; i++ {
            selected = append(selected, i)
            backtrack(i+1, k)
            selected = selected[:len(selected)-1]
        }
    }
    backtrack(1, k)
    return result
}