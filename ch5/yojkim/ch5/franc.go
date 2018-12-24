package ch5

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	franc := Franc{}
	franc.amount = amount

	return franc
}

func (f Franc) Times(multiplier int) *Franc {
	return &Franc{f.amount * multiplier}
}

func (f Franc) Equal(input Dollar) bool {
	return f.amount == input.amount
}
