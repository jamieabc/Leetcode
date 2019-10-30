package problem_65

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	assert.Equal(t, true, isNumber("0"), "wrong")
	assert.Equal(t, true, isNumber("3."), "wrong")
	assert.Equal(t, true, isNumber("+.8"), "wrong")
	assert.Equal(t, true, isNumber(" 005047e+6"), "wrong")
	assert.Equal(t, false, isNumber("-."), "wrong")
	assert.Equal(t, false, isNumber(" "), "wrong")
	assert.Equal(t, false, isNumber("0e"), "wrong")
	assert.Equal(t, false, isNumber(".e1"), "wrong")
}
