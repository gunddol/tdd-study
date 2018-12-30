package ch7

type Money interface {
	Times(multiplier int) *Money
	Equal(input Money) bool
}