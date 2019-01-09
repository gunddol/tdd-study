package ch13

type Expression interface {
	Reduce(to string) Money
}