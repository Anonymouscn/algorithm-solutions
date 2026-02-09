func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    l, r := 0, m * n
    for l < r {
        mid := (l + r) / 2
        if v := matrix[mid/n][mid%n]; v == target {
            return true
        } else if v > target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return false
}