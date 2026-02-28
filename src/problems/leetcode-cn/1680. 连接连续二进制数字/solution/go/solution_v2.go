func concatenatedBinary(n int) int {
    result, mod := int64(0), int64(1_000_000_007)
    for i := 1; i <= n; i++ {
        w := bits.Len(uint(i))
        result = (result << w | int64(i)) % mod
    }
    return int(result)
}