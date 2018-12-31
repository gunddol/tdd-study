package ch9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.Times(2))
	assert.Equal(t, NewFranc(15), five.Times(3))
}

func TestFrancEquality(t *testing.T) {
	assert.True(t, NewFranc(5).Equal(NewFranc(5)))
	assert.False(t, NewFranc(5).Equal(NewFranc(6)))
	assert.False(t, NewFranc(5).Equal(NewDollar(5)))
}