package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	assert.Equal(t, 10, product.Amount)
	product = five.Times(3)
	assert.Equal(t, 15, product.Amount)
}
