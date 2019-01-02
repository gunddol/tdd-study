package ch4;

import org.junit.Test;
import static org.junit.Assert.*;

public class dollar_Test {

    @Test
    public void testMultiplication() {
        dollar five = new dollar(5);
        assertEquals(new dollar(10), five.times(2));
        assertEquals(new dollar(15), five.times(3));
    }

    @Test
    public void testEquality(){
        assertTrue(new dollar(5).equals(new dollar(5)));
        assertFalse(new dollar(5).equals(new dollar(6)));
    }
}