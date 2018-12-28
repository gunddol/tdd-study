package ch4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, *five.Times(2), NewDollar(10))
	assert.Equal(t, *five.Times(3), NewDollar(15))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equal(NewDollar(5)))
	assert.False(t, NewDollar(5).Equal(NewDollar(6)))
}