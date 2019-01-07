package ch14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(5).Currency())
	assert.NotEqual(t, "CHF", NewDollar(5).Currency())
}

func TestMultiplication(t *testing.T) {
	dollar := NewDollar(5)
	assert.Equal(t, NewDollar(10), dollar.Times(2))
	assert.Equal(t, NewDollar(15), dollar.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equal(NewDollar(5)))
	assert.False(t, NewDollar(5).Equal(NewDollar(6)))
	assert.False(t, NewDollar(5).Equal(NewFranc(5)))
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	res := five.Plus(five)
	sum := res.(Sum)
	assert.Equal(t, five, sum.Augend)
	assert.Equal(t, five, sum.Addend)
}