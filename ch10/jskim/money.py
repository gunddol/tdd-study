import unittest
from abc import *

class Money():

    _amount:int
    _currency:str

    def __init__(self, amount:int, currency:str):
        self._amount = amount
        self._currency = currency

    @classmethod
    def dollar(self, amount:int):
        return Dollar(amount, "USD")
    
    @classmethod
    def franc(self, amount:int):
        return Franc(amount, "CHF")

    def times(self, multiplier:int):
        return Money(self._amount * multiplier, self._currency)
    
    def currency(self):
        return self._currency

    def equals(self, money:object):
        return (self._amount == money._amount) & (self.currency() == money.currency())

class Dollar(Money):

    def __init__(self, amount:int, currency:str):
        super().__init__(amount, currency)

class Franc(Money):

    def __init__(self, amount:int, currency:str):
        super().__init__(amount, currency)

class testDifferentClassEqualilty(unittest.TestCase):

    def testDifferentClassEquality(self):
        self.assertTrue(Money(10, "CHF").equals(Franc(10, "CHF")))

class testCurrency(unittest.TestCase):
    
    def testCurrency(self):
        self.assertEqual("USD", Money.dollar(1).currency())

        self.assertEqual("CHF", Money.franc(1).currency())


class testEquality(unittest.TestCase):
    
    def testEquality(self):
        self.assertTrue(Money.dollar(5).equals(Money.dollar(5)))

        self.assertFalse(Money.dollar(5).equals(Money.dollar(6)))

        self.assertTrue(Money.franc(5).equals(Money.franc(5)))

        self.assertFalse(Money.franc(5).equals(Money.franc(6)))

        self.assertFalse(Money.franc(5).equals(Money.dollar(5)))

class testMultiplication(unittest.TestCase):

    def testMultiplication(self):
        five = Money.dollar(5)

        self.assertTrue(Money.dollar(10).equals(five.times(2)))

        self.assertTrue(Money.dollar(15).equals(five.times(3)))

class testFrancMultiplication(unittest.TestCase):

    def testFrancMultiplication(self):
        five = Money.franc(5)

        self.assertTrue(Money.franc(10).equals(five.times(2)))

        self.assertTrue(Money.franc(15).equals(five.times(3)))
    
if __name__ == "__main__":
    unittest.main()