package function

func CheckInNumberSlice[T uint64 | int32](a T, s []T) bool {
	for _, val := range s {
		if a == val {
			return true
		}
	}
	return false
}

func DelEleInSlice[T uint64 | int32](a T, old []T) []T {
	for i, val := range old {
		if val == a {
			return append(old[:i], old[i+1:]...)
		}
	}
	return old
}
