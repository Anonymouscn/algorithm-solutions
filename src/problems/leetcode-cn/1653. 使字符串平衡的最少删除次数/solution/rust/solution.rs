impl Solution {
    pub fn minimum_deletions(s: String) -> i32 {
        let length = s.len();
        if length < 2 {
            return 0;
        }

        let (mut bc, mut del) = (0, 0);

        for v in s.bytes() {
            if v == b'b' {
                bc += 1;
            } else {
                del += 1;
                if del > bc {
                    del = bc
                }
            }
        }

        del
    }
}