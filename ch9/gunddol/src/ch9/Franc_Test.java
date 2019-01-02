package ch9;

import org.junit.Test;

import static org.junit.Assert.*;

public class Franc_Test {

    @Test
    public void testFrancMultiplication() {
        Money five = Money.franc(5);
        assertEquals(Money.franc(10), five.times(2));
        assertEquals(Money.franc(15), five.times(3));
    }
}