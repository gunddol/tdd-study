package ch8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestDollarEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equal(NewDollar(5)))
	assert.False(t, NewDollar(5).Equal(NewDollar(6)))
	assert.False(t, NewDollar(5).Equal(NewFranc(5)))
}