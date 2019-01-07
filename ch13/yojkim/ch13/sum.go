package ch13

type Sum struct {
	Augend Money
	Addend Money
}

func NewSum(augend Money, addend Money) Sum {
	return Sum{
		Augend: augend,
		Addend: addend,
	}
}

func (s Sum) Reduce(to string) Money {
	amount := s.Augend.amount + s.Addend.amount
	return NewMoney(amount, to)
}