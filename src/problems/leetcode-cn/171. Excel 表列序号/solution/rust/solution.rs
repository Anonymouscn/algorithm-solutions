impl Solution {
    pub fn title_to_number(column_title: String) -> i32 {
        let mut res = 0;
        for b in column_title.bytes() {
            res = res*26 + ((b - 'A' as u8) as i32 + 1);
        }
        res
    }
}