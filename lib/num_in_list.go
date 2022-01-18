package lib

func NumInList(list []int, num int) bool {
	for _, idx := range list {
		if idx == num {
			return true
		}
	}
	return false
}
