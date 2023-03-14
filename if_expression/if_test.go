package if_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturn(t *testing.T) {
	r := Return(true, "是", "否")
	assert.Equal(t, r, "是")

	f := Return(false, "是", "否")
	assert.Equal(t, f, "否")
}

func TestReturnByFunc(t *testing.T) {
	r := ReturnByFunc(true, func() string {
		return "是"
	}, func() string {
		return "否"
	})
	assert.Equal(t, r, "是")
	assert.NotEqual(t, r, "否")
}
