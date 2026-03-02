func addDigits(num int) int {
    t := num
    for t / 10 > 0 {
        v, tv := 0, t
        for ; tv > 0; tv /= 10 {
            v += tv % 10
        }
        t = v
    }
    return t
}