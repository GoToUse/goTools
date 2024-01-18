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
