package reversedstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
