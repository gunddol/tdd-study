package ch8;

import org.junit.Test;

import static org.junit.Assert.*;

public class Money_Test {
    @Test
    public void testEquality(){
        assertTrue(Money.dollar(5).equals(Money.dollar(5)));
        assertFalse(Money.dollar(5).equals(Money.dollar(6)));
        assertTrue(Money.franc(5).equals(Money.franc(5)));
        assertFalse(Money.franc(5).equals(Money.franc(6)));
        assertFalse(Money.franc(5).equals(Money.dollar(5))); //금액과 클래스가 모두 같아야 참
    }
}