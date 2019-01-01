import unittest

class Dollar:
    
    amount:int

    def __init__(self, amount:int):
        self.amount = amount

    def times(self, multiplier:int):
        self.amount *= multiplier

class testMultiplication(unittest.TestCase):

    def testMultiplication(self):
        five = Dollar(5)
        five.times(2)

        self.assertEqual(10, five.amount)
    
if __name__ == "__main__":
    unittest.main()