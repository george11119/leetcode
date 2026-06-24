import unittest

from k_closest_points_to_origin import Solution


class TestKthLargest(unittest.TestCase):
    def test_basic_sequence(self):
        s = Solution()
        print(s.kClosest([[0, 2], [2, 2]], 1))
        print(s.kClosest([[0, 2], [2, 0], [2, 2]], 2))


if __name__ == "__main__":
    unittest.main()
