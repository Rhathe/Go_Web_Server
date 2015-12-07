package gson

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {
	assert := assert.New(t)

	a := Gson{
		"foo": "bar",
		"baz": Gson{
			"bee": "boo",
			"baa": []int{1, 2, 3},
		},
	}

	b := map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			"bee": "boo",
			"baa": []int{1, 2, 3},
		},
	}

	c := map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			"bee": "boo",
			"baa": []int{1, 2, 3, 4},
		},
	}

	d := map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			"beea": "boo",
			"baa": []int{1, 2, 3},
		},
	}

	assert.Equal(a.IsEqual(b), true, "they should be equal")
	assert.Equal(a.IsEqual(c), false, "they shouldn't be equal")
	assert.Equal(a.IsEqual(d), false, "they shouldn't be equal")
}
