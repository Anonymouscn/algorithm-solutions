class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        result = 0
        for c in columnTitle.encode('utf-8'):
            result = result * 26 + (c - ord('A') + 1)
        return result