func rotateString(s string, goal string) bool {
	l1, l2 := len(s), len(goal)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if s == goal {
			return true
		}
		s = s[1:] + string(s[0])
	}
	return false
}