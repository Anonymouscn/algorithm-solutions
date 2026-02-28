func concatenatedBinary(n int) int {
    result, mod, bl := int64(0), int64(1_000_000_007), 0
    for i := 1; i <= n; i++ {
        if i & (i-1) == 0 {
            bl++
        }
        result = (result << bl + int64(i)) % mod
    }
    return int(result)
}