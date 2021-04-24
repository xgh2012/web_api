package utils

//查找是否存在
func InArray(value int, check []int) bool {
	for _, val := range check {
		if val == value {
			return true
		}
	}
	return false
}
