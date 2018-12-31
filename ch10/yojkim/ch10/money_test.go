package ch10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(5).Currency())
	assert.Equal(t, "CHF", NewFranc(5).Currency())

	assert.NotEqual(t, "CHF", NewDollar(5).Currency())
	assert.NotEqual(t, "USD", NewFranc(5).Currency())
}