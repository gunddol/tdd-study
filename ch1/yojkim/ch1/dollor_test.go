package ch1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)
	assert.Equal(t, 10, five.Amount)
}
