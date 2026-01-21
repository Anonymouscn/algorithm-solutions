func plusOne(digits []int) []int {
    l := len(digits)
    result := make([]int, l + 1)
    digits[l-1]++
    r := 0
    for i := l-1; i >= 0; i-- {
        digits[i] += r
        if digits[i] > 9 {
            r = digits[i] / 10
            digits[i] = digits[i] % 10
        } else {
            r = 0
        }
        result[i+1] = digits[i]
    }
    if r > 0 {
        result[0] = r
    } else {
        result = result[1:]
    }
    return result
}