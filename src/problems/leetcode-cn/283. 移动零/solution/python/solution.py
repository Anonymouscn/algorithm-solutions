class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        length = len(nums)
        if length < 2:
            return
        slow, fast = 0, 0
        while fast < length:
            if nums[fast] != 0:
                if fast != slow:
                    nums[fast], nums[slow] = nums[slow], nums[fast]
                slow += 1
            fast += 1