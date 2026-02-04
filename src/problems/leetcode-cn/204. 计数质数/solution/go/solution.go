// 素数筛(埃氏筛) [O(NloglogN)] / 线性筛 [接近O(n)], 示例使用线性筛法解决
func countPrimes(n int) int {
    primes, isComp := make([]int, 0), make([]bool, n)
    for i := 2; i < n; i++ {
        if !isComp[i-1] {
            primes = append(primes, i)
        }

        for _, p := range(primes) {
            v := i * p

            if v >= n {
                break
            }

            isComp[v-1] = true

            if i % p == 0 {
                break
            }
        }
    }
    return len(primes)
}