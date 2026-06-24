import unittest
from last_stone_weight import Solution


class TestLastStoneWeight(unittest.TestCase):
    def test(self):
        tt = [
            {"stones": [2, 3, 6, 2, 4], "want": 1},
            {"stones": [1, 2], "want": 1}
        ]

        for tc in tt:
            s = Solution()
            got = s.lastStoneWeight(tc["stones"])
            want = tc["want"]

            self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
