import unittest

from kth_largest_element_in_array import Solution


class TestKthLargestInArray(unittest.TestCase):
    def test_basic_sequence(self):
        s = Solution()
        self.assertEqual(s.findKthLargest([2, 3, 1, 5, 4], 2), 4)
        self.assertEqual(s.findKthLargest([2, 3, 1, 1, 5, 5, 4], 3), 4)


if __name__ == "__main__":
    unittest.main()
