import unittest
from car_pooling import Solution


class TestGetOrder(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example1(self):
        self.assertEqual(self.solution.carPooling([[4, 1, 2], [3, 2, 4]], 4), True)

    def test_example2(self):
        self.assertEqual(self.solution.carPooling([[2, 1, 3], [3, 2, 4]], 4), False)


if __name__ == "__main__":
    unittest.main()
