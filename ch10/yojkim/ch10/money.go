package ch10

type Money interface {
	Times(multiplier int) Money
	Equal(input Money) bool
	Currency() string
}