// 版本一：回溯 + 状态备份剪枝 O(mn)
func minPathSum(grid [][]int) int {
    rows, cols := len(grid), 0
    if rows > 0 {
        cols = len(grid[0])
    }

    backup := make(map[uint64]int)
    var hash func(a, b int) uint64
    hash = func(a, b int) uint64 {
        return uint64(uint32(a))<<32 | uint64(uint32(b))
    }

    var minPath func(point []int) int
    minPath = func(point []int) int {
        k := hash(point[0], point[1])
        if v, ok := backup[k]; ok {
            return v
        }

        mask := math.MaxInt / 4
        b, r := mask, mask
        if point[0]+1 < rows {
            b = minPath([]int{point[0]+1, point[1]})
        }
        if point[1]+1 < cols {
            r = minPath([]int{point[0], point[1]+1})
        }
        if b == mask && r == mask {
            backup[k] = grid[point[0]][point[1]]
            return backup[k]
        }
        backup[k] = min(b, r) + grid[point[0]][point[1]]
        return backup[k]
    }
    
    return minPath([]int{0, 0})
}