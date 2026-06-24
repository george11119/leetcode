import unittest

from task_scheduler import Solution


class TestKthLargestInArray(unittest.TestCase):
    def test_basic_sequence(self):
        s = Solution()
        self.assertEqual(s.leastInterval(["X", "X", "Y", "Y"], 2), 5)
        self.assertEqual(s.leastInterval(["A", "A", "A", "B", "C"], 3), 9)


if __name__ == "__main__":
    unittest.main()
