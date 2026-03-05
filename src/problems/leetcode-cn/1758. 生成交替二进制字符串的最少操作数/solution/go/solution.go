func minOperations(s string) int {
    flag, length, c1, c2 := byte(0), len(s), 0, 0
    for i := 0; i < length; i++ {
        if v := s[i]; v != flag + '0' {
            c1++
        } else {
            c2++
        }
        flag ^= 1
    }
    return min(c1, c2)
}