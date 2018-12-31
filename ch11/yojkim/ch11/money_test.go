package ch11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewMoney(5, "USD").Currency())
	assert.NotEqual(t, "CHF", NewMoney(5, "USD").Currency())
}

func TestMultiplication(t *testing.T) {
	dollar := NewMoney(5, "USD")
	assert.Equal(t, NewMoney(10, "USD"), dollar.Times(2))
	assert.Equal(t, NewMoney(15, "USD"), dollar.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewMoney(5, "USD").Equal(NewMoney(5, "USD")))
	assert.False(t, NewMoney(5, "USD").Equal(NewMoney(6, "USD")))
	assert.False(t, NewMoney(5, "USD").Equal(NewMoney(5, "CHF")))
}