func combinationSum(candidates []int, target int) [][]int {
    result, selected := make([][]int, 0), make([]int, 0)
    sort.Ints(candidates) // 题目"无重复元素", 可排序剪枝
    var backtrack func(start int, remain int) 
    backtrack = func(start int, remain int) {
        if remain == 0 {
            tmp := make([]int, len(selected))
            copy(tmp, selected)
            result = append(result, tmp)
            return
        }
        for i := start; i < len(candidates); i++ {
            v := candidates[i]
            if v > remain { // 剪枝缩小搜索范围
                break
            }
            selected = append(selected, v)
            backtrack(i, remain - v)
            selected = selected[:len(selected)-1]
        }
    }
    backtrack(0, target)
    return result
}