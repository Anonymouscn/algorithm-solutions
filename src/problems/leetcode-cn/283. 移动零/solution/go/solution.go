func moveZeroes(nums []int)  {
    length := len(nums)
    if length < 2 {
        return
    }

    slow, fast := 0, 0
    for ; fast < length; fast++ {
        if nums[fast] != 0 {
            if slow != fast {
                nums[slow], nums[fast] = nums[fast], nums[slow]
            }
            slow++
        }
    }
}