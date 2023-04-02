package bytetostring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToString(t *testing.T) {
	r := "hello, bytetostring"
	r_b := StringToBytes(r)
	conv := BytesToString(r_b)
	assert.Equal(t, r, conv)
}
