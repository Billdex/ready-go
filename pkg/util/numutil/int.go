package numutil

func IntIf(condition bool, a int, b int) int {
	if condition {
		return a
	}
	return b
}

// IntLeast Int 数字 n 必须至少为某数 lower，否则返回条件下界 lower
func IntLeast(n, lower int) int {
	return IntIf(n >= lower, n, lower)
}

// IntMost Int 数字 n 至多为某数 upper，否则返回条件上界 upper
func IntMost(n, upper int) int {
	return IntIf(n <= upper, n, upper)
}

// IntBetween Int 数字 n 必须在区间 [lower, upper] 内，否则返回其最接近的上界或下界
func IntBetween(n, lower, upper int) int {
	return IntMost(IntLeast(n, lower), upper)
}
