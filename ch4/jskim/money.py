import unittest

class Dollar:
    
    __amount:int

    def __init__(self, amount:int):
        self.__amount = amount

    def times(self, multiplier:int):
        return Dollar(self.__amount * multiplier)
    
    def equals(self, dollar:object):
        return self.__amount == dollar.__amount

class testEquality(unittest.TestCase):
    
    def testEquality(self):
        self.assertTrue(Dollar(5).equals(Dollar(5)))

        self.assertFalse(Dollar(5).equals(Dollar(6)))

class testMultiplication(unittest.TestCase):

    def testMultiplication(self):
        five = Dollar(5)

        self.assertTrue(Dollar(10).equals(five.times(2)))

        self.assertTrue(Dollar(15).equals(five.times(3)))
    
if __name__ == "__main__":
    unittest.main()