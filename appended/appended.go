package appended

func Merge[T comparable](arrays ...[]T) []T {
	var c []T
	for _, a := range arrays {
		c = append(c, a...)
	}
	return c
}

func ConcatCopyPreAllocate[T comparable](slices [][]T) []T {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]T, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}
