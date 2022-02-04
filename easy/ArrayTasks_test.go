package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"

)

func Test_twoSums(t *testing.T) {
	nums := []int{3,2,4}
	target := 6

	actual := twoSums(nums, target)
	expected := []int{1,2}
	assert.Equal(t, expected, actual)
}
