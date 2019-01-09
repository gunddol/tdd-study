package ch15;

class Money implements Expression{

    protected String currency;
    protected int amount;

    Expression times(int multiplier){
        return new Money(amount * multiplier, currency);
    }

    Money(int amount, String currency){
        this.amount = amount;
        this.currency = currency;
    }

    public static Money dollar(int amount){
        return new Money(amount, "USD");
    }

    public static Money franc(int amount){
        return new Money(amount,"CHF");
    }

    public String toString(){
        return amount + " " + currency;
    }

    String currency(){
        return currency;
    }

    public boolean equals(Object object){
        Money money = (Money) object;
        return amount == money.amount
                && currency().equals(money.currency());
    }

    public Expression plus(Expression addend) {
        return new Sum(this, addend);
    }

    public Money reduce(Bank bank, String to){
        int rate = bank.rate(currency, to);
        return new Money(amount / rate, to);
    }
}
