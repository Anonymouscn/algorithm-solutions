func maxDotProduct(nums1 []int, nums2 []int) int {
    rows, cols := len(nums2), len(nums1)
    dpm := make([][]int, rows)
    // 计算点积矩阵
    for r := 0; r < rows; r++ {
        arr := make([]int, cols)
        for c := 0; c < cols; c++ {
            arr[c] = nums2[r] * nums1[c]
        }
        dpm[r] = arr
    }
    // 二维 DP
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            cmp := make([]int, 0)
            cmp = append(cmp, dpm[r][c])
            if r > 0 {
                cmp = append(cmp, dpm[r-1][c])
            }
            if c > 0 {
                cmp = append(cmp, dpm[r][c-1])
            }
            if r > 0 && c > 0 {
                cmp = append(cmp, dpm[r-1][c-1])
                cmp = append(cmp, dpm[r-1][c-1] + dpm[r][c])
            }
            dpm[r][c] = max(cmp)
        }
    }
    return dpm[rows-1][cols-1]
}

func max(arr []int) int {
    length := len(arr)
    if length == 0 {
        return -1
    }
    max := arr[0]
    for i := 1; i < length; i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}