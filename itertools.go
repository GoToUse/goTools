// itertools.go is simply like python's itertools library
// to furnish many convenient and powerful tools to
// help us to develop more fast.

package goTools

import (
	"strings"
)

// PairwiseCombine refer to python's `itertools.combinations`.
func PairwiseCombine(conditions []string, r int) []string {
	n := len(conditions)
	if r > n {
		return nil
	}
	indices := generateArrayFrom(0, r)

	result := make([]string, 0, n)

	result = append(result, getItem(indices, conditions))

outer:
	for {
		var _tmp int
		for _, i := range reverseArray(generateArrayFrom(0, r)) {
			_tmp = i
			if indices[i] != i+n-r {
				break
			}
			if i == 0 {
				break outer
			}
		}
		indices[_tmp] += 1

		for _, j := range generateArrayFrom(_tmp+1, r) {
			indices[j] = indices[j-1] + 1
		}
		result = append(result, getItem(indices, conditions))
	}

	return result
}

// getItem should support more output form. TODO:
func getItem(idx []int, array []string) string {
	combineArray := make([]string, 0, len(array))
	for _, i := range idx {
		combineArray = append(combineArray, array[i])
	}
	return strings.Join(combineArray, ",")
}

func generateArrayFrom(s, e int) []int {
	_a := make([]int, 0, e)
	for i := s; i < e; i++ {
		_a = append(_a, i)
	}

	return _a
}

func reverseArray[T comparable](a []T) []T {
	i := 0
	j := len(a) - 1

	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	return a
}
