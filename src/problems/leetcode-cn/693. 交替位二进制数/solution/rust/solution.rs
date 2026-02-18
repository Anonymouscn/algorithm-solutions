impl Solution {
    pub fn has_alternating_bits(mut n: i32) -> bool {
        n = n ^ (n >> 1);
        n & (n + 1) == 0
    }
}