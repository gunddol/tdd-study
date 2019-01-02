package ch4;

public class dollar {
    private int amount;

    public dollar(int amount){
        this.amount = amount;
    }

    public dollar times(int multiplier) {
        return new dollar(amount * multiplier);
    }

    public boolean equals(Object object){
        dollar dollar = (dollar) object;
        return amount == dollar.amount;
    }
}
