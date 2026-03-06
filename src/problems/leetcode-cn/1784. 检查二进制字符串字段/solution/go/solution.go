func checkOnesSegment(s string) bool {
    l, r, length, cnt := 0, 0, len(s), 0
    for l < length && r < length {
        if s[l] == '1' {
            for r = l + 1; r < length && s[r] == '1'; r++ {}
            cnt++
            if cnt > 1 {
                return false
            }
            l = r
        }
        l++
    }
    return true
}