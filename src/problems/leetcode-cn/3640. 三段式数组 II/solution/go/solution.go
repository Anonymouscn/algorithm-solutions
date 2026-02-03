func maxSumTrionic(nums []int) int64 {
    mask := int64(math.MinInt64 / 4)
    length, result := len(nums), mask

    // 计算差分数组
    diff := make([]int, length)
    for i := 1; i < length; i++ {
        diff[i] = nums[i] - nums[i-1]
    }

    // 获取所有最大连续递增区间 [start, end]
    p1, p2, up := 1, 1, make([][]int, 0)
    for p1 < length && p2 < length {
        if diff[p1] > 0 {
            for p2 = p1; p2 + 1 < length && diff[p2+1] > 0; p2++ {}
            up = append(up, []int{p1-1, p2})
            p1 = p2 + 1
        } else {
            p1++
        }
        p2 = p1
    }

    // 计算总和
    p1, p2, length = 0, 1, len(up)
    for p1 < length && p2 < length {
        // 排除平坦区间 (非连续递减区间无法构成[增减增]三段式)
        if is_flat(diff, up[p1][1], up[p2][0]) {
            p1++
            p2++
            continue
        }

        s1 := int64(nums[up[p1][1]] + nums[up[p1][1]-1])
        s2 := int64(nums[up[p2][0]] + nums[up[p2][0]+1])
        s3 := int64(0)
        if up[p2][0] - up[p1][1] > 1 {
            s3 = sum(nums, up[p1][1] + 1, up[p2][0] - 1)
        }

        // 计算最优单调范围
        if up[p1][1] - up[p1][0] > 1 {
            st := s1
            for i := up[p1][1]-2; i >= up[p1][0]; i-- {
                st += int64(nums[i])
                if st > s1 {
                    s1 = st
                }
            }
        }
        if up[p2][1] - up[p2][0] > 1 {
            st := s2
            for i := up[p2][0]+2; i <= up[p2][1]; i++ {
                st += int64(nums[i])
                if st > s2 {
                    s2 = st
                }
            }
        }

        if total := s1 + s2 + s3; total > result {
            result = total
        }

        p1++
        p2++
    }

    return result
}

func sum(arr []int, a, b int) int64 {
    sum := int64(0)
    for i := a; i <= b; i++ {
        sum += int64(arr[i])
    }
    return sum
}

func is_flat(diff []int, a, b int) bool {
    ptr, end := a + 1, b + 1
    for ; ptr < end; ptr++ {
        if diff[ptr] == 0 {
            return true
        }
    }
    return false
}