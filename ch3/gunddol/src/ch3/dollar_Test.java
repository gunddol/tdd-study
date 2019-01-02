package ch3;

import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;


public class dollar_Test {

    @Test
    public void testMultiplication() {
        dollar five = new dollar(5);
        dollar product = five.times(2);
        assertEquals(10, product.amount);
        product = five.times(3);
        assertEquals(15, product.amount);
    }

    @Test
    public void testEquality(){
        assertTrue(new dollar(5).equals(new dollar(5)));
        assertFalse(new dollar(5).equals(new dollar(6)));
    }
}