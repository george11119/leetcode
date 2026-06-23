import unittest

from kth_largest_integer import KthLargest


class TestKthLargest(unittest.TestCase):
    def test_basic_sequence(self):
        kth = KthLargest(3, [4, 5, 8, 2])
        self.assertEqual(kth.add(3), 4)
        self.assertEqual(kth.add(5), 5)
        self.assertEqual(kth.add(10), 5)
        self.assertEqual(kth.add(9), 8)
        self.assertEqual(kth.add(4), 8)

    def test_empty_initial_nums(self):
        kth = KthLargest(1, [])
        self.assertEqual(kth.add(-3), -3)
        self.assertEqual(kth.add(-2), -2)
        self.assertEqual(kth.add(-4), -2)
        self.assertEqual(kth.add(0), 0)
        self.assertEqual(kth.add(4), 4)

    def test_k_equals_one(self):
        kth = KthLargest(1, [5])
        self.assertEqual(kth.add(10), 10)
        self.assertEqual(kth.add(7), 10)


if __name__ == "__main__":
    unittest.main()
