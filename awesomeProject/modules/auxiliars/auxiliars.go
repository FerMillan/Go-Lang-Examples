package auxiliars

func Index[T comparable](s []T, find T) int {
	for i, v := range s {
		if v == find {
			return i
		}
	}
	return -1
}

func IsExist[T comparable](s []T, find T) bool {
	for _, v := range s {
		if v == find {
			return true
		}
	}
	return false
}
