func reverse(x int) int {
    res, scope := 0, []int{-1 << 31, (1 << 31) - 1}
    for x != 0 {
        v := x % 10
        if res*10 + v > scope[1] || res*10 + v < scope[0] {
            return 0
        }
        res = res * 10 + v
        x /= 10
    }
    return res
}