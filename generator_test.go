package goTools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_Generate(t *testing.T) {
	g := NewGenerator(1, 10, 2)
	g.Generate()

	var res []int
	for i := 0; i < 3; i++ {
		fmt.Println("Iteration:", i+1)
		for val, ok := g.Next(); ok; val, ok = g.Next() {
			res = append(res, val)
		}
		assert.Equal(t, []int{1, 3, 5, 7, 9}, res)

		// 重置生成器并重新开始
		g.Reset()
		g.Generate()
		res = []int{}
	}
	g.Close()
}
