// itertools.go is simply like python's itertools library
// to furnish many convenient and powerful tools to
// help us to develop more fast.

package goTools

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func factorial(n int) int {
	var b int = 1

	for n > 0 {
		b *= n
		n--
	}

	return b
}

// Combinations refer to python's `itertools.combinations`.
func Combinations[T any](conditions []T, r int) [][]T {
	n := len(conditions)
	if r > n {
		return nil
	}
	indices := rangeFrom(0, r)

	result := make([][]T, 0, factorial(n)/factorial(r))

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

// Repeat refer to python's `itertools.Repeat`.
// Make an iterator that returns object over and over again.
// Runs indefinitely unless the times argument is greater than 0.
// Repeat(10, 3) --> 10 10 10
func Repeat[T any, R constraints.Integer](v T, times R) chan T {
	ch := make(chan T)

	switch {
	case times > 0:
		go func() {
			defer close(ch)
			for times > 0 {
				times--
				ch <- v
			}
		}()
	default:
		go func() {
			defer close(ch)
			for {
				ch <- v
			}
		}()
	}

	return ch
}

// Count refer to python's `itertools.Count`.
// Make an iterator that returns evenly spaced values starting with number start.
// Count(10) --> 10 11 12 13 14 ...
// Count(2.5, 0.5) --> 2.5 3.0 3.5 ...
func Count[T constraints.Integer | constraints.Float](start T, times T) chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for {
			ch <- start
			start += times
		}
	}()

	return ch
}

// Operate 策略模式
type Operate[T constraints.Float | constraints.Integer] interface {
	// execute is used to implement various operator protocol.
	execute(total, element T) T
}

// Add +
type Add[T constraints.Float | constraints.Integer] struct {
}

func (Add[T]) execute(total, element T) T {
	return total + element
}

// Sub -
type Sub[T constraints.Float | constraints.Integer] struct {
}

func (Sub[T]) execute(total, element T) T {
	return total - element
}

// Mul *
type Mul[T constraints.Float | constraints.Integer] struct {
}

func (Mul[T]) execute(total, element T) T {
	return total * element
}

// Div /
type Div[T constraints.Float | constraints.Integer] struct {
}

func (Div[T]) execute(total, element T) T {
	return total / element
}

// Accumulate refer to python's `itertools.Accumulate`.
// Make an iterator that returns accumulated sums,
// or accumulated results of other binary functions (specified via the optional func argument).
// By now it only supports four kind of operators.
func Accumulate[T constraints.Float | constraints.Integer](iter []T, operator string, initial T) []T {
	var operatorFunc Operate[T]

	switch operator {
	case "-":
		operatorFunc = Sub[T]{}
	case "/":
		operatorFunc = Div[T]{}
	case "*":
		operatorFunc = Mul[T]{}
	case "+":
		operatorFunc = Add[T]{}
	default:
		err := fmt.Errorf("invalid operator: %s", operator)
		fmt.Println(err.Error())
		return nil
	}

	total := initial

	result := make([]T, 0, len(iter)+1)

	result = append(result, total)
	for _, element := range iter {
		total = operatorFunc.execute(total, element)
		result = append(result, total)
	}

	return result
}

// Pairwise refer to python's `itertools.Pairwise`.
// Return successive overlapping pairs taken from the input iterable.
func Pairwise[T constraints.Ordered](iter []T) [][]T {
	result := make([][]T, 0, len(iter))

	for i, j := 0, 1; j < len(iter); i, j = i+1, j+1 {
		result = append(result, []T{iter[i], iter[j]})
	}

	return result
}

// Chain refer to python's `itertools.Chain`.
// Make an iterator that returns elements from the first iterable until it is exhausted,
// then proceeds to the next iterable, until all of the iterables are exhausted.
// Used for treating consecutive sequences as a single sequence.
func Chain[T constraints.Ordered](iter ...[]T) []T {
	if len(iter) == 1 {
		return iter[0]
	}

	if len(iter) == 0 {
		return []T{}
	}

	capacity := 0
	for _, item := range iter {
		capacity += len(item)
	}

	result := make([]T, 0, capacity)
	for _, item := range iter {
		result = append(result, item...)
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
