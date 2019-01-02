package ch8;

public class dollar extends Money{

    public dollar(int amount){
        this.amount = amount;
    }

    public Money times(int multiplier) {
        return new dollar(amount * multiplier);
    }

}
