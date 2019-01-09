package ch16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Plus(fiveBucks)
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(15), result)
}

func TestSumTimes(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Times(2)
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(20), result)
}