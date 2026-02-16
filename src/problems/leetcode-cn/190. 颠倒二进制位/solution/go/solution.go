func reverseBits(n int) int {
    seq := make([]byte, 0)
    for n > 0 {
        seq = append(seq, byte(n % 2))
        n >>= 1
    }
    for len(seq) < 32 {
        seq = append(seq, 0)
    }
    sum, p := 0, 1
    for i := 0; i < 32; i++ {
        sum += int(seq[32 - i - 1]) * p
        p <<= 1
    }
    return sum
}