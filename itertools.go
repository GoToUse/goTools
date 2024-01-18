// itertools.go is simply like python's itertools library
// to furnish many convenient and powerful tools to
// help us to develop more fast.

package goTools

import (
	"golang.org/x/exp/constraints"
)

// Combinations refer to python's `itertools.combinations`.
func Combinations[T any](conditions []T, r int) [][]T {
	n := len(conditions)
	if r > n {
		return nil
	}
	indices := rangeFrom(0, r)

	result := make([][]T, 0, n)

	result = append(result, toSlice(indices, conditions))

outer:
	for {
		var _tmp int
		for _, i := range reverseArray(rangeFrom(0, r)) {
			_tmp = i
			if indices[i] != i+n-r {
				break
			}
			if i == 0 {
				break outer
			}
		}
		indices[_tmp] += 1

		for _, j := range rangeFrom(_tmp+1, r) {
			indices[j] = indices[j-1] + 1
		}
		result = append(result, toSlice(indices, conditions))
	}

	return result
}

// toSlice should support more output form. TODO:
func toSlice[T any](idx []int, array []T) []T {
	combineArray := make([]T, 0, len(array))
	for _, i := range idx {
		combineArray = append(combineArray, array[i])
	}
	return combineArray
}

// rangeFrom like range in python, default step is 1.
func rangeFrom[T constraints.Integer](s, e T) []T {
	_a := make([]T, 0, e-s)
	for i := s; i < e; i++ {
		_a = append(_a, i)
	}

	return _a
}

// reverseArray like reverse in python.
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

// reverseArray like reverse in python.
func reverseArrayNoChange[T comparable](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)

	return reverseArray(b)
}
