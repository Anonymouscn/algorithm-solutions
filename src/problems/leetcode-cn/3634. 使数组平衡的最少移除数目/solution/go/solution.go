func minRemoval(nums []int, k int) int {
    length := len(nums)
    if length < 2 {
        return 0
    }

    sort.Ints(nums)

    // 滑动窗口
    l, r, inc := 0, 1, 0
    for r < length {
        if nums[r] <= nums[l]*k {
            inc = max(inc, r-l+1)
            r++
        } else {
            l++
        }
    }

    return length - inc
}