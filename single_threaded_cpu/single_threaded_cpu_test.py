import unittest
from single_threaded_cpu import Solution


class TestGetOrder(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example1(self):
        self.assertEqual(self.solution.getOrder([[1, 4], [3, 3], [2, 1]]), [0, 2, 1])

    def test_example2(self):
        self.assertEqual(
            self.solution.getOrder([[1, 2], [2, 4], [3, 2], [4, 1]]),
            [0, 2, 3, 1],
        )


if __name__ == "__main__":
    unittest.main()
