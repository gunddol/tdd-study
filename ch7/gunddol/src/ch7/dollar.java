package ch7;

public class dollar extends Money{

    public dollar(int amount){
        this.amount = amount;
    }

    public dollar times(int multiplier) {
        return new dollar(amount * multiplier);
    }

    public boolean equals(Object object){
        Money money = (Money) object;
        return amount == money.amount;
    }
}
