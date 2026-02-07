class Solution:
    def minimumDeletions(self, s: str) -> int:
        length = len(s)
        if length < 2:
            return 0

        bc, delete = 0, 0
        
        for v in s:
            if v == 'b':
                bc += 1
            else:
                delete += 1
                if delete > bc:
                    delete = bc

        return delete