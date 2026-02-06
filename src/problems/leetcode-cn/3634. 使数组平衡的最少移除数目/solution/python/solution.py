class Solution:
    def minRemoval(self, nums: List[int], k: int) -> int:
        length = len(nums)
        if length < 2:
            return 0

        nums.sort()
        l, r, inc = 0, 1, 0

        while r < length:
            if nums[r] <= nums[l]*k:
                inc = max(inc, r-l+1)
                r += 1
            else:
                l += 1
        
        return length-inc