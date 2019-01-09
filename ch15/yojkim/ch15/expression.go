package ch15

type Expression interface {
	Reduce(bank Bank, to string) Money
	Plus(addend Expression) Expression
}