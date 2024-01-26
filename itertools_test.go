package goTools

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	r1 := Combinations([]string{"A", "B", "C", "D"}, 2)
	assert.Equal(t, [][]string{{"A", "B"}, {"A", "C"}, {"A", "D"}, {"B", "C"}, {"B", "D"}, {"C", "D"}}, r1)

	r2 := Combinations([]int{0, 1, 2, 3}, 3)
	assert.Equal(t, [][]int{{0, 1, 2}, {0, 1, 3}, {0, 2, 3}, {1, 2, 3}}, r2)

	r3 := Combinations([]bool{true, false, false}, 2)
	assert.Equal(t, [][]bool{{true, false}, {true, false}, {false, false}}, r3)
}

func TestRangeFrom(t *testing.T) {
	r := rangeFrom(-4, -1)
	assert.Equal(t, []int{-4, -3, -2}, r)

	rr := lo.RangeFrom(-8, 7)
	assert.Equal(t, []int{-8, -7, -6, -5, -4, -3, -2}, rr)

	rrr := lo.RangeWithSteps(-8, -1, 1)
	assert.Equal(t, []int{-8, -7, -6, -5, -4, -3, -2}, rrr)
}

func TestReverseSlice(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := reverseArrayNoChange(a)

	assert.NotEqual(t, a, b)
}

func TestRepeat(t *testing.T) {
	result := make([]string, 0, 8)
	for res := range Repeat("a", 5) {
		result = append(result, res)
	}
	assert.Equal(t, []string{"a", "a", "a", "a", "a"}, result)
}

func TestCount(t *testing.T) {
	threshold := 10000

	loops := 0
	for v := range Count(1000, 10) {
		loops++
		if v == threshold {
			break
		}
	}

	assert.Equal(t, 901, loops)
}

func TestAccumulate(t *testing.T) {
	res := Accumulate([]float64{1, 2, 3, 4, 5}, "&", 2)
	assert.Equal(t, []float64(nil), res)

	res1 := Accumulate([]int8{1, 2, 3, 4, 5}, "*", 2)
	assert.Equal(t, []int8{2, 2, 4, 12, 48, -16}, res1)

	res2 := Accumulate([]int16{1, 2, 3, 4, 5}, "*", 2)
	assert.Equal(t, []int16{2, 2, 4, 12, 48, 240}, res2)

	res3 := Accumulate([]int16{1, 2, 3, 4, 5}, "+", 2)
	assert.Equal(t, []int16{2, 3, 5, 8, 12, 17}, res3)

	res4 := Accumulate([]int16{1, 2, 3, 4, 5}, "-", 2)
	assert.Equal(t, []int16{2, 1, -1, -4, -8, -13}, res4)

	res5 := Accumulate([]int16{1, 2, 3, 4, 5}, "/", 2)
	assert.Equal(t, []int16{2, 2, 1, 0, 0, 0}, res5)
}

func TestPairwise(t *testing.T) {
	res := Pairwise([]int{2, 3, 4, 9, 10})
	assert.Equal(t, [][]int{{2, 3}, {3, 4}, {4, 9}, {9, 10}}, res)

	res1 := Pairwise([]string{"A", "B", "C", "D"})
	assert.Equal(t, [][]string{{"A", "B"}, {"B", "C"}, {"C", "D"}}, res1)
}

func TestChain(t *testing.T) {
	res := Chain([]string{"a", "b", "c"}, []string{"d", "e", "f", "g"}, []string{"h", "i", "j", "k", "l", "m", "n"})
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}, res)

	res1 := Chain([]int{1, 2, 3}, []int{6, 9, 10}, []int{})
	assert.Equal(t, []int{1, 2, 3, 6, 9, 10}, res1)
}
