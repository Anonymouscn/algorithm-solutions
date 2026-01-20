func minBitwiseArray(nums []int) []int {
    length := len(nums)
    for i := 0; i < length; i++ {
        if nums[i] & 1 == 0 { // 偶数不可能成立，最低位必为1，必为奇数
            nums[i] = -1
        } else {
            // 位运算: 在二进制连续的1里，清除掉低位连续1中最高位的1，实现 y | (y + 1) = target，同时y最小（2进制转换对应10进制的数最小）
            n := nums[i] + 1
            nums[i] ^= ((n & -n) >> 1)
        }
    }
    return nums
}