func minimumDeletions(s string) int {
    length := len(s)
    if length < 2 {
        return 0
    }

    bc, del := 0, 0 // [这次待删除b的次数, 上次最小操作次数]

    for i := 0; i < length; i++ {
        if s[i] == 'a' {
            // del = min(del+1, bc)
            del++
            if del > bc {
                del = bc
            }
        } else {
            bc++
        }
    }

    return del
}