class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        m = {}
        for i, v in enumerate(nums):
            t = target - v
            if t in m:
                return [m[t], i]
            else:
                m[v] = i
        return [0, 0]