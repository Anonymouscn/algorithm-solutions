func minFlips(s string) int {
    // 初始化窗口
    odd, even, window := []int{0, 0}, []int{0, 0}, len(s)
    min_ops := window
    for i := 0; i < window; i++ {
        if i % 2 == 0 {
            odd[int(s[i] - '0')]++
        } else {
            even[int(s[i] - '0')]++
        }
    }
    min_ops = min(odd[1] + even[0], odd[0] + even[1])

    // 偶数项不需要滑动窗口改变模式
    if window % 2 == 0 {
        return min_ops
    }

    // 奇数项滑动窗口找全局最小值
    length := 2 * window
    for i := window; i < length; i++ {
        idx := i % window
        odd[int(s[idx] - '0')]--
        even[int(s[idx] - '0')]++
        odd, even = even, odd
        min_ops = min(min_ops, min(odd[1] + even[0], odd[0] + even[1]))
    }

    return min_ops
}