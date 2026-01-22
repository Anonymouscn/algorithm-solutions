func generateParenthesis(n int) []string {
    result := []string{}

    var backtrack func(cur string, open, close, max int)

    backtrack = func(cur string, open, close, max int) {
        if len(cur) == max*2 {
            result = append(result, cur)
        }
        if open < max {
            backtrack(cur+"(", open+1, close, max)
        }
        if close < open {
            backtrack(cur+")", open, close+1, max)
        }
    }

    backtrack("", 0, 0, n)

    return result
}