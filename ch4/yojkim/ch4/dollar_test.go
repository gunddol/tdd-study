package ch4

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), *five.Times(2))
	assert.Equal(t, NewDollar(15), *five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5).Equal(NewDollar(5)), true)
	assert.Equal(t, NewDollar(5).Equal(NewDollar(6)), false)
}