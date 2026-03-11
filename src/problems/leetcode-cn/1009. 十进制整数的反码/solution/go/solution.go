func bitwiseComplement(n int) int {
    if n == 0 {
        return 1
    }
    arr, ans := make([]int, 0), 0
    for ; n > 0; n >>= 1 {
        arr = append(arr, n & 1)
    }
    p := 1
    for i := 0; i < len(arr); i++ {
        ans += (arr[i] ^ 1) * p
        p <<= 1
    }
    return ans
}