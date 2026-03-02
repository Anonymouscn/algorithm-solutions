func findKthBit(n int, k int) byte {
    result := []byte{0}
    for i := 1; i < n; i++ {
        length := len(result)
        t := make([]byte, length)
        for i := 0; i < length; i++ {
            t[i] = result[length-1-i] ^ 1
        }
        result = append(append(result, 1), t...)
    }
    return '0' + result[k-1]
}