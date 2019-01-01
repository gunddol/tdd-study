import unittest
from abc import *

class Money(metaclass=ABCMeta):

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

    @abstractmethod
    def times(self, multiplier:int):
        pass
    
    def currency(self):
        return self._currency

    def equals(self, money:object):
        return (self._amount == money._amount) & (self.__class__.__name__ == money.__class__.__name__)

class Dollar(Money):

    def __init__(self, amount:int, currency:str):
        super().__init__(amount, currency)

    def times(self, multiplier:int):
        return Money.dollar(self._amount * multiplier)

class Franc(Money):

    def __init__(self, amount:int, currency:str):
        super().__init__(amount, currency)

    def times(self, multiplier:int):
        return Money.franc(self._amount * multiplier)

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