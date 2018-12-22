package ch1

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)
	assert.Equal(t, five.Amount, 10)
}
