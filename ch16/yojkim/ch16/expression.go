package ch16

type Expression interface {
	Reduce(bank Bank, to string) Money
	Plus(addend Expression) Expression
	Times(multiplier int) Expression
}