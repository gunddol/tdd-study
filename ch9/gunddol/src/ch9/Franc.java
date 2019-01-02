package ch9;

public class Franc extends Money{

    protected String currency;

    public Franc(int amount, String currency){
        super(amount, currency);
    }

    Money times(int multiplier) {
        return Money.franc(amount * multiplier);
    }

//    String currency(){
//        return currency;
//    }
}
