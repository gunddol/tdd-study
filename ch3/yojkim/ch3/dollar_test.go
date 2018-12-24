package ch3

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	assert.Equal(t, product.Amount, 10)
	product = five.Times(3)
	assert.Equal(t, product.Amount, 15)
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5).Equal(NewDollar(5)), true)
	assert.Equal(t, NewDollar(5).Equal(NewDollar(6)), false)
}