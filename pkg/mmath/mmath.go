package mmath

// MinInt 比较两个整数的大小返回最小的
func MinInt(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
