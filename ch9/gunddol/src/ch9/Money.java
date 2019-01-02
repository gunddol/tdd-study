package ch9;

abstract class Money {
    abstract Money times(int multiplier);

    protected String currency;
    protected int amount;

    Money(int amount, String currency){
        this.amount = amount;
        this.currency = currency;
    }

    public static Money dollar(int amount){
        return new dollar(amount, "USD");
    }

    public static Money franc(int amount){
        return new Franc(amount,"CHF");
    }

    String currency(){
        return currency;
    }

    public boolean equals(Object object){
        Money money = (Money) object;
        return amount == money.amount
                && getClass().equals(money.getClass());
    }
}
