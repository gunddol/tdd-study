package ch2;

public class dollar {
    int amount;

    public dollar(int amount){
        this.amount = amount;
    }

    public dollar times(int multiplier) {
        return new dollar(amount * multiplier);
    }
}
