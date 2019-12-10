package unit_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Pet struct {
	name string
}

func TestSomething(t *testing.T) {
	assert.Equal(t, 1, 1, "1 equal 1")
	assert.NotEqual(t, 1, 2, "1 not equal 2")
	// assert for nil (good for errors)
	var dog Pet
	t.Logf("%p", &dog)
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, dog) {
		// now we know that dog isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "", dog.name)
	}
}
