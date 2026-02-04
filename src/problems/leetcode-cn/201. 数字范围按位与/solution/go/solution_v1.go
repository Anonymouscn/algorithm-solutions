func rangeBitwiseAnd(left int, right int) int {
    shitf := 0
    for left < right {
        left >>= 1
        right >>= 1
        shitf++
    }
    return left << shitf
}