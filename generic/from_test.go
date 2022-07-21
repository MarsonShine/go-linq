package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromQuery(t *testing.T) {
	intArray := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrFilted := From(intArray).Where(func(i int32) bool { return i%2 == 0 }).ToSlice()
	assert.NotEqual(t, 0, len(arrFilted))
	assert.Equal(t, 5, len(arrFilted))
	assert.Equal(t, int32(2), arrFilted[0])
}
