package ch5;

import org.junit.Test;

import static org.junit.Assert.*;

public class Franc_Test {

    @Test
    public void testFrancMultiplication() {
        Franc five = new Franc(5);
        assertEquals(new Franc(10), five.times(2));
        assertEquals(new Franc(15), five.times(3));
    }

}