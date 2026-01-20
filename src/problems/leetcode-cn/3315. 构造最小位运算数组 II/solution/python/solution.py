class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        for i, e in enumerate(nums):
            if (e & 1 == 0):
                nums[i] = -1
            else:
                n = nums[i] + 1
                nums[i] ^= ((n & -n) >> 1)
        return nums