import unittest

class Money:
    
    _amount:int

    def equals(self, money:object):
        return (self._amount == money._amount) & (self.__class__.__name__ == money.__class__.__name__)

class Dollar(Money):

    def __init__(self, amount:int):
        self._amount = amount

    def times(self, multiplier:int):
        return Dollar(self._amount * multiplier)

class Franc(Money):

    def __init__(self, amount:int):
        self._amount = amount

    def times(self, multiplier:int):
        return Franc(self._amount * multiplier)

class testEquality(unittest.TestCase):
    
    def testEquality(self):
        self.assertTrue(Dollar(5).equals(Dollar(5)))

        self.assertFalse(Dollar(5).equals(Dollar(6)))

        self.assertTrue(Franc(5).equals(Franc(5)))

        self.assertFalse(Franc(5).equals(Franc(6)))

        self.assertFalse(Franc(5).equals(Dollar(5)))

class testMultiplication(unittest.TestCase):

    def testMultiplication(self):
        five = Dollar(5)

        self.assertTrue(Dollar(10).equals(five.times(2)))

        self.assertTrue(Dollar(15).equals(five.times(3)))

class testFrancMultiplication(unittest.TestCase):

    def testFrancMultiplication(self):
        five = Franc(5)

        self.assertTrue(Franc(10).equals(five.times(2)))

        self.assertTrue(Franc(15).equals(five.times(3)))
    
if __name__ == "__main__":
    unittest.main()