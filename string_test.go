package goTools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToString(t *testing.T) {
	r := "Hello, byteToString"
	rb := StringToBytes(r)
	conv := BytesToString(rb)
	assert.Equal(t, r, conv)
}

func TestReverseString(t *testing.T) {
	s := "Hello world"
	assert.Equal(t, "dlrow olleH", reverseString(s))
}

func TestReverseStringByByte(t *testing.T) {
	s := "Hello world!"
	assert.Equal(t, "!dlrow olleH", reverseStringByByte(s))
}

func TestReverseStringByEmptyStr(t *testing.T) {
	s := ":Hello world!!!"
	assert.Equal(t, "!!!dlrow olleH:", reverseStringByEmptyStr(s))
}
