func wordPattern(pattern string, s string) bool {
    arr, m1, m2, length := strings.Split(s, " "), make(map[byte]string), make(map[string]int), 0
    if length = len(arr); length != len(pattern) {
        return false
    }
    for i := 0; i < length; i++ {
        if v, ok := m1[pattern[i]]; ok {
            if arr[i] != v {
                return false
            }
        } else {
            if _, ok := m2[arr[i]]; ok {
                return false
            }
            m1[pattern[i]] = arr[i]
            m2[arr[i]] = i
        }
    }
    return true
}