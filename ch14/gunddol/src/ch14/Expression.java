package ch14;

public interface Expression {
    Money reduce(Bank bank, String to);
}
