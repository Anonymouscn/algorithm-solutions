func constructTransformedArray(nums []int) []int {
    length := len(nums)
    result := make([]int, length)
    // 向左走 x 位 <=> 向右走 length - x 位
    for i, v := range(nums) {
        result[i] = nums[(length+(i+v)%length)%length]
    }
    return result
}