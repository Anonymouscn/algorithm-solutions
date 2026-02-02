func isTrionic(nums []int) bool {
    start, length := 0, len(nums)

    // 少于 4 个不能构成完整坡道
    if length < 4 {
        return false
    }

    // 爬坡闭包函数 (模拟)
    climb := func(cmp func(a, b int) bool) bool {
        origin, end := start, length - 1
        for ; start < end && cmp(nums[start+1], nums[start]); start++ {}
        return start > origin
    }

    if !climb(func(a, b int) bool { return a > b }) ||
    !climb(func(a, b int) bool { return a < b }) ||
    !climb(func(a, b int) bool { return a > b }) {
        return false
    }

    return start == length-1
}