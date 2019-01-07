package ch14

type Money struct {
	currency string
	amount int
}

func NewMoney(amount int, currency string) Money {
	return Money{
		amount: amount,
		currency: currency,
	}
}

func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount * multiplier, m.currency)
}

func (m Money) Equal(input Money) bool {
	if m.currency != input.currency {
		return false
	}
	return m.amount == input.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Plus(addend Money) Expression {
	return NewSum(m, addend)
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount / rate, to)
}