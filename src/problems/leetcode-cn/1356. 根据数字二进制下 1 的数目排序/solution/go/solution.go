func sortByBits(arr []int) []int {
    m, keys, result := make(map[int][]int), []int{}, []int{}
    for i := 0; i < len(arr); i++ {
        k := count(arr[i])
        if v, ok := m[k]; ok {
            v = append(v, arr[i])
            m[k] = v
        } else {
            keys = append(keys, k)
            m[k] = []int{arr[i]}
        }
    }
    sort.Ints(keys)
    for i := 0; i < len(keys); i++ {
        v := m[keys[i]]
        if len(v) > 1 {
            sort.Ints(v)
        }
        result = append(result, v...)
    }
    return result
}

func count(i int) int {
    cnt := 0
    for ; i > 0; i >>= 1 {
        cnt += (i & 1)
    }
    return cnt
}