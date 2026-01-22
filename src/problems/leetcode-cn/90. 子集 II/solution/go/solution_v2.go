func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    result, m := make([][]int, 0), make(map[string]struct{})
    backtrack([]int{}, nums, func(node []int) {
        hash := fmt.Sprintf("%v", node)
        if _, ok := m[hash]; !ok {
            result = append(result, append([]int(nil), node...))
            m[hash] = struct{}{}
        }
    })
    return result
}

func backtrack(selected, available []int, onSelect func(node []int)) {
    onSelect(selected)
    for i := 0; i < len(available); i++ {
        backtrack(append(selected, available[i]), available[i+1:], onSelect)
    }
}