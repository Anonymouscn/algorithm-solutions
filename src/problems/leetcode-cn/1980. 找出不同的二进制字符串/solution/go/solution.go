func findDifferentBinaryString(nums []string) string {
    n, m := len(nums), make(map[string]int)
    for i := 0; i < n; i++ {
        m[nums[i]] = 1
    }

    arr := make([]byte, 0)
    for i := 0; i < n; i++ {
        arr = append(arr, '0')
    }

    if _, ok := m[string(arr)]; !ok {
        return string(arr)
    }

    for {
        success, v := addOne(arr)
        if !success {
            break
        }
        if _, ok := m[v]; !ok {
            return v
        }
    }

    return ""
}

func addOne(arr []byte) (success bool, ans string) {
    n, remain := len(arr), 1
    for i := n - 1; i >= 0 && remain > 0; i-- {
        current := int(arr[i] - '0') + remain
        arr[i] = '0' + byte(current & 1)
        remain = current / 2

        if remain == 0 {
            success = true
            ans = string(arr)
            return
        }
    }
    return
}