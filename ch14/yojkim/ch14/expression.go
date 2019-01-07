package ch14

type Expression interface {
	Reduce(bank Bank, to string) Money
}