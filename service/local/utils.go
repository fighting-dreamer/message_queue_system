package service


func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}


func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
