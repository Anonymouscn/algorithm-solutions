func readBinaryWatch(turnedOn int) []string {
    result, m := []string{}, make(map[string]int)

    var backtrack func(hours, mins []int, hour, min, limit int)
    backtrack = func(hours, mins []int, hour, min, limit int) {
        if limit == 0 {
            m[fmt.Sprintf("%d:%02d", hour, min)] = 1
        }

        for i, v := range hours {
            if t := hour + v; t < 12 {
                nh := []int{}
                if i+1 < len(hours) {
                    nh = append(nh, hours[i+1:]...)
                }
                backtrack(nh, mins, t, min, limit-1)
            }
        }
        for i, v := range mins {
            if t := min + v; t < 60 {
                nm := []int{}
                if i+1 < len(mins) {
                    nm = append(nm, mins[i+1:]...)
                }
                backtrack(hours, nm, hour, t, limit-1)
            }
        }
    }

    // hours[1, 2, 4, 8], mins[1, 2, 4, 8, 16, 32]
    backtrack([]int{1, 2, 4, 8}, []int{1, 2, 4, 8, 16, 32}, 0, 0, turnedOn)

    for v, _ := range m {
        result = append(result, v)
    }
    return result
}