import unittest

class Dollar:
    
    amount:int

    def __init__(self, amount:int):
        self.amount = amount

    def times(self, multiplier:int):
        return Dollar(self.amount * multiplier)
    
    def equals(self, dollar:object):
        return self.amount == dollar.amount

class testEquality(unittest.TestCase):
    
    def testEquality(self):
        self.assertTrue(Dollar(5).equals(Dollar(5)))

        self.assertFalse(Dollar(5).equals(Dollar(6)))

class testMultiplication(unittest.TestCase):

    def testMultiplication(self):
        five = Dollar(5)
        product = five.times(2)

        self.assertEqual(10, product.amount)

        product = five.times(3)

        self.assertEqual(15, product.amount)
    
if __name__ == "__main__":
    unittest.main()