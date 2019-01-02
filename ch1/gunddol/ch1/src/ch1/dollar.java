package ch1;

public class dollar {
    int amount;

    public dollar(int amount){
        this.amount = amount;
    }

    public void times(int multiplier){
        amount *= multiplier;
    }
}
