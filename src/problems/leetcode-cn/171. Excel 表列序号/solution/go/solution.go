// A = 0*26+1
// AB = (0*26+1)*26 + 2 = 28
// BA = (0*26+2)*26 + 1 = 53
// ZY = (0*26+26)*26 + 25 = 701
// AAA = ((0*26+1)*26+1)*26 + 1 = 703
func titleToNumber(columnTitle string) int {
    res := 0
    for _, v := range([]rune(columnTitle)) {
        res = res*26 + int(v - 'A' + 1)
    }
    return res
}