package ch8

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount}
}

func (f Franc) Times(multiplier int) Money {
	return Franc{f.amount * multiplier}
}

func (f Franc) Equal(input Money) bool {
	franc, ok := input.(Franc)
	if !ok {
		return false
	}
	return f.amount == franc.amount
}
