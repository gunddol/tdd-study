package ch7

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	franc := Franc{amount}
	return franc
}

func (f Franc) Times(multiplier int) *Franc {
	return &Franc{f.amount * multiplier}
}

func (f Franc) Equal(input Franc) bool {
	return f.amount == input.amount
}
