func minimumPairRemoval(nums []int) int {
    ok, length, cnt := accept(nums), len(nums), 0
    if length < 2 || ok {
        return cnt
    }
    for !ok {
        cnt++
        i, _ := min(nums)
        nums = merge(nums, i-1, i)
        ok = accept(nums)
    }
    return cnt
}

func merge(nums []int, i1, i2 int) []int {
    length := len(nums)
    nums[i1] += nums[i2]
    for i := i2; i < length - 1; i++ {
        nums[i] = nums[i+1]
    }
    nums = nums[:length-1]
    return nums
}

func min(nums []int) (int, int) {
    length := len(nums)
    if length < 2 {
        return -1, -1
    }
    m, mv := 1, nums[1] + nums[0]
    for i := 2; i < length; i++ {
        if t := nums[i] + nums[i-1]; t < mv {
            m = i
            mv = t
        }
    }
    return m, mv
}

func accept(nums []int) bool {
    length := len(nums)
    for i := 1; i < length; i++ {
        if nums[i] < nums[i-1] {
            return false
        }
    }
    return true
}