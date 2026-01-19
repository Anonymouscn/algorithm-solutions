class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        result = []
        for v in nums:
            result += [-1]
            for t in range(0, v):
                if (t | (t + 1)) == v:
                    result[len(result)-1] = t
                    break
        return result