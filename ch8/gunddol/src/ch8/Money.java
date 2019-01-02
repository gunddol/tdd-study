package ch8;

abstract class Money {
    abstract Money times(int multiplier);
    protected int amount;

    public boolean equals(Object object){
        Money money = (Money) object;
        return amount == money.amount
                && getClass().equals(money.getClass());
    }

    public static Money dollar(int amount){
        return new dollar(amount);
    }

    public static Money franc(int amount){
        return new Franc(amount);
    }
}
