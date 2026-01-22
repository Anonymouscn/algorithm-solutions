func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    result, path := [][]int{}, []int{}

    var backtrack func(n int)

    backtrack = func(n int) {
        if n == len(nums) {
            result = append(result, append([]int(nil), path...))
            return
        }

        i := n
        for ; i < len(nums) && nums[i] == nums[n]; i++ {}
        cnt := i - n

        backtrack(i)

        for t := 0; t < cnt; t++ {
            path = append(path, nums[n])
            backtrack(i)
        }

        path = path[:len(path)-cnt]
    }

    backtrack(0)

    return result
}