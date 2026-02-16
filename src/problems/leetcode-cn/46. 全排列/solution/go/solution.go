func permute(nums []int) [][]int {
    result := [][]int{}

    var backtrack func(selected, available []int)
    backtrack = func(selected, available []int) {
        if len(available) == 0 {
            result = append(result, selected)
        }
        for i, v := range available {
            se, av := make([]int, len(selected)), make([]int, i)
            copy(se, selected)
            copy(av, available[:i])
            se, av = append(se, v), append(av, available[i+1:]...)
            backtrack(se, av)
        }
    }

    backtrack([]int{}, nums)

    return result
}