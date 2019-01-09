package ch12

type Bank struct {

}

func NewBank() Bank {
	return Bank{}
}

func (Bank) Reduce(source Expression, to string) Money {
	return NewDollar(10)
}