package ch1;

import org.junit.Test;

import static org.junit.Assert.*;

public class dollar_Test {

    @Test
    public void testMultiplication(){
        dollar five =  new dollar(5);
        five.times(2);
        assertEquals(10, five.amount);
    }
}