package ch9;

import org.junit.Test;
import static org.junit.Assert.*;

public class dollar_Test {

    @Test
    public void testMultiplication() {
        Money five = Money.dollar(5);
        assertEquals(Money.dollar(10), five.times(2));
        assertEquals(Money.dollar(15), five.times(3));
    }

}