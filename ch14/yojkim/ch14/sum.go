package ch14

type Sum struct {
	Augend Money
	Addend Money
}

func NewSum(augend, addend Money) Sum {
	return Sum{
		Augend: augend,
		Addend: addend,
	}
}

func (s Sum) Reduce(bank Bank, to string) Money {
	amount := s.Augend.amount + s.Addend.amount
	return NewMoney(amount, to)
}