package ch2;

import org.junit.Test;

import static org.junit.Assert.*;

public class dollar_Test {
    @Test
    public void testMultiplication() {
        dollar five = new dollar(5);
        dollar product = five.times(2);
        assertEquals(10, product.amount);
        product = five.times(3);
        assertEquals(15, product.amount);
    }
}