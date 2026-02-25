func numSteps(s string) int {
    cnt, bs := 0, []byte(s)
    for length := len(bs); length > 0; length = len(bs) {
        if bs[length-1] == '0' {
            bs = bs[:length-1]
        } else {
            if length == 1 {
                break
            }

            for remain, p := 1, length-1; remain > 0 && p >= 0; p-- {
                if bs[p] == '0' {
                    bs[p] = '1'
                    break
                } else {
                    bs[p] = '0'
                }
            }
        }
        cnt++
    }
    return cnt
}