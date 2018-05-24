package twitter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wipeinc/wipeinc/twitter"
)

var minusIDListTests = []struct {
	a        []int64
	b        []int64
	expected []int64
}{
	{
		[]int64{},
		[]int64{},
		[]int64{},
	},
	{
		[]int64{1, 2, 3},
		[]int64{},
		[]int64{1, 2, 3},
	},
	{
		[]int64{},
		[]int64{1, 2, 3},
		[]int64{},
	},
	{
		[]int64{1, 2, 3},
		[]int64{4, 5, 6},
		[]int64{1, 2, 3},
	},
	{
		[]int64{1, 2, 3, 4},
		[]int64{2, 4, 5, 6},
		[]int64{1, 3},
	},
}

func TestMinusIDList(t *testing.T) {
	var result []int64
	assert := assert.New(t)
	for _, tt := range minusIDListTests {
		result = twitter.MinusIDList(tt.a, tt.b)
		assert.Equal(tt.expected, result)
	}
}
