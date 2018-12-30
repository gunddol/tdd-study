package ch7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), *five.Times(2))
	assert.Equal(t, NewDollar(15), *five.Times(3))
}

func TestDollarEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5).Equal(NewDollar(5)), true)
	assert.Equal(t, NewDollar(5).Equal(NewDollar(6)), false)
}