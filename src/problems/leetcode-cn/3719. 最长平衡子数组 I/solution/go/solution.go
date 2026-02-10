func longestBalanced(nums []int) int {
    length, result := len(nums), 0

    var backtrack func(l, r, odd, even int, m map[int]int) int
    backtrack = func(l, r, odd, even int, m map[int]int) int {
        cnt := 0
        if _, ok := m[nums[r]]; !ok {
            m[nums[r]] = 1
            if nums[r] % 2 == 0 {
                even++
            } else {
                odd++
            }
        } else {
            m[nums[r]]++
        }
        if odd == even {
            cnt = r+1-l
        }
        if next := r+1; next < length {
            cnt = max(cnt, backtrack(l, next, odd, even, m))
        }
        return cnt
    }

    for i := 0; i < length; i++ {
        m := make(map[int]int)
        result = max(result, backtrack(i, i, 0, 0, m))
    }

    return result
}