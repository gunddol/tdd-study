package ch10

type Franc struct {
	amount int
	currency string
}

func NewFranc(amount int) Franc {
	return Franc{
		amount: amount,
		currency: "CHF",
	}
}

func (f Franc) Times(multiplier int) Money {
	return NewFranc(f.amount * multiplier)
}

func (f Franc) Equal(input Money) bool {
	franc, ok := input.(Franc)
	if !ok {
		return false
	}
	return f.amount == franc.amount
}

func (f Franc) Currency() string {
	return f.currency
}