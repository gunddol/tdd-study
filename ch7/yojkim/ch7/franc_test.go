package ch7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), *five.Times(2))
	assert.Equal(t, NewFranc(15), *five.Times(3))
}