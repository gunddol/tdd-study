package ch9;

public class dollar extends Money{

    public String currency;

    public dollar(int amount, String currency){
        super(amount, currency);
    }

    Money times(int multiplier) {
        return Money.dollar(amount * multiplier);
    }

//    String currency(){
//        return currency;
//    }
}
