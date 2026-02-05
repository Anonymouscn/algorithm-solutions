class Solution:
    def constructTransformedArray(self, nums: List[int]) -> List[int]:
        result, l = [], len(nums)
        for i, v in enumerate(nums):
            result.append(nums[((v + i) % l + l) % l])
        return result