import unittest

from twitter import Twitter


class TestTwitter(unittest.TestCase):
    def test_example_sequence(self):
        twitter = Twitter()

        self.assertIsNone(twitter.postTweet(1, 10))
        self.assertIsNone(twitter.postTweet(2, 20))

        self.assertEqual(twitter.getNewsFeed(1), [10])
        self.assertEqual(twitter.getNewsFeed(2), [20])

        self.assertIsNone(twitter.follow(1, 2))
        #
        self.assertEqual(twitter.getNewsFeed(1), [20, 10])
        self.assertEqual(twitter.getNewsFeed(2), [20])

        self.assertIsNone(twitter.unfollow(1, 2))

        self.assertEqual(twitter.getNewsFeed(1), [10])

    def test_follow_unfollow_feed(self):
        twitter = Twitter()

        twitter.postTweet(1, 101)
        twitter.postTweet(1, 102)
        twitter.follow(2, 1)

        self.assertEqual(twitter.getNewsFeed(2), [102, 101])

        twitter.unfollow(2, 1)

        self.assertEqual(twitter.getNewsFeed(2), [])


if __name__ == "__main__":
    unittest.main()
