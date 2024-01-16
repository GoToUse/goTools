package goTools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairwiseCombine(t *testing.T) {
	r := PairwiseCombine([]string{"A", "B", "C", "D"}, 2)
	assert.Equal(t, []string{"A,B", "A,C", "A,D", "B,C", "B,D", "C,D"}, r)
}
