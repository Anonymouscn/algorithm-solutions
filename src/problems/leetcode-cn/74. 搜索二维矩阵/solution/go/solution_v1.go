func searchMatrix(matrix [][]int, target int) bool {
    t, b, l, r := 0, len(matrix), 0, len(matrix[0])-1
    if target < matrix[t][l] || target > matrix[b-1][r] {
        return false
    }

    search_row, search_col := t, l

    for t < b {
        search_row = (t + b) / 2
        if matrix[search_row][0] == target {
            return true
        }
        if search_row == t || search_row == b {
            break
        }
        if matrix[search_row][0] > target {
            b = search_row
        } else if matrix[search_row][0] < target {
            t = search_row
        }
    }

    for l <= r {
        search_col = (l + r) / 2
        if matrix[search_row][search_col] == target {
            return true
        }
        if matrix[search_row][search_col] > target {
            r = search_col - 1
        } else if matrix[search_row][search_col] < target {
            l = search_col + 1
        }
    }

    return false
}