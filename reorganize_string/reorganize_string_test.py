import unittest
from reorganize_string import Solution


class TestGetOrder(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example1(self):
        self.assertEqual(len(self.solution.reorganizeString("axyy")), 4)

    def test_example2(self):
        self.assertEqual(self.solution.reorganizeString("ccccd"), "")

    def test_example3(self):
        self.assertEqual(self.solution.reorganizeString("aaabb"), "ababa")


if __name__ == "__main__":
    unittest.main()
