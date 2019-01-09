import unittest
from abc import *

class Expression():
    
    @abstractmethod
    def plus(self, addend:object):
        pass

    @abstractmethod
    def reduce(self, bank:object, to:str):
        pass

class Bank():

    __rates = {}

    def reduce(self, source:object, to:str):
        return source.reduce(self, to)
    
    def rate(self, where:str, to:str):
        if where == to:
            return 1
        
        rate = self.__rates[Pair(where, to).keyString()]

        return int(rate)
    
    def addRate(self, where:str, to:str, rate:int):
        self.__rates[Pair(where, to).keyString()] = rate

class Pair():

    __where:str
    __to:str

    def __init__(self, where:str, to:str):
        self.__where = where
        self.__to = to
    
    def keyString(self):
        return str(self.__where + ',' + self.__to)

    def equals(self, pair:object):
        return self.__where.equals(pair.__where) & self.__to.equals(pair.__to)

    def hashCode(self):
        return 0

class Money(Expression):

    _amount:int
    _currency:str

    def __init__(self, amount:int, currency:str):
        self._amount = amount
        self._currency = currency

    @classmethod
    def dollar(self, amount:int):
        return Money(amount, "USD")
    
    @classmethod
    def franc(self, amount:int):
        return Money(amount, "CHF")

    def times(self, multiplier:int):
        return Money(self._amount * multiplier, self._currency)
    
    def plus(self, addend:object):
        return Summation(self, addend)

    def reduce(self, bank:object, to:str):
        rate = bank.rate(self._currency, to)

        return Money(self._amount / rate, to)
    
    def currency(self):
        return self._currency

    def equals(self, money:object):
        return (self._amount == money._amount) & (self.currency() == money.currency())

class Summation(Expression):
    
    augend:object
    addend:object

    def __init__(self, augend:object, addend:object):
        self.augend = augend
        self.addend = addend
    
    def reduce(self, bank:object, to:str):
        return Money(self.augend._amount + self.addend._amount, to)

class testIdentityRate(unittest.TestCase):

    def testIdentityRate(self):
        self.assertEqual(1, Bank().rate("USD", "USD"))

class testReduceMoneyDifferentCurrency(unittest.TestCase):

    def testReduceMoneyDifferentCurrency(self):
        bank = Bank()
        bank.addRate("CHF", "USD", 2)
        result = bank.reduce(Money.franc(2), "USD")

        self.assertTrue(Money.dollar(1).equals(result))

class testReduceMoney(unittest.TestCase):

    def testReduceMoney(self):
        bank = Bank()
        result = bank.reduce(Money.dollar(1), "USD")

        self.assertTrue(Money.dollar(1).equals(result))

class testReduceSum(unittest.TestCase):

    def testReduceSum(self):
        summation = Summation(Money.dollar(3), Money.dollar(4))
        bank = Bank()
        result = bank.reduce(summation, "USD")

        self.assertTrue(Money.dollar(7).equals(result))

class testPlusReturnsSum(unittest.TestCase):

    def testPlusReturnsSum(self):
        five = Money.dollar(5)
        summation = five.plus(five)

        self.assertTrue(five.equals(summation.augend))
        self.assertTrue(five.equals(summation.addend))

class testSimpleAddition(unittest.TestCase):

    def testSimpleAddition(self):
        five = Money.dollar(5)
        summation = five.plus(five)
        bank = Bank()
        summation = Money.dollar(5).plus(Money.dollar(5))
        reduced = bank.reduce(summation, "USD")

        self.assertTrue(Money.dollar(10).equals(reduced))

class testCurrency(unittest.TestCase):
    
    def testCurrency(self):
        self.assertEqual("USD", Money.dollar(1).currency())

        self.assertEqual("CHF", Money.franc(1).currency())


class testEquality(unittest.TestCase):
    
    def testEquality(self):
        self.assertTrue(Money.dollar(5).equals(Money.dollar(5)))

        self.assertFalse(Money.dollar(5).equals(Money.dollar(6)))

        self.assertFalse(Money.franc(5).equals(Money.dollar(5)))

class testMultiplication(unittest.TestCase):

    def testMultiplication(self):
        five = Money.dollar(5)

        self.assertTrue(Money.dollar(10).equals(five.times(2)))

        self.assertTrue(Money.dollar(15).equals(five.times(3)))
    
if __name__ == "__main__":
    unittest.main()