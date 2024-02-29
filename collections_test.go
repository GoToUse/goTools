package goTools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCounter(t *testing.T) {
	data := []int{1, 1, 2, 3, 5, 8, 10, 3}

	c := NewCounter(data)
	assert.Equal(t, map[any]int{1: 2, 3: 2, 5: 1, 2: 1, 8: 1, 10: 1}, c.data)

	data1 := "abceaefdb"
	c1 := NewCounter(data1)
	assert.Equal(t, map[any]int{"a": 2, "b": 2, "c": 1, "e": 2, "f": 1, "d": 1}, c1.data)

	data2 := map[any]int{"apple": 3, "banana": 2}
	c2 := NewCounter(data2)
	assert.Equal(t, map[any]int{"apple": 3, "banana": 2}, c2.data)
}
