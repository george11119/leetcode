import heapq
from collections import defaultdict


class Twitter:
    def __init__(self):
        self.tweets = defaultdict(list)
        self.follow_list = defaultdict(set)
        self.count = 0

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.tweets[userId].append([self.count, tweetId])
        self.count += 1

    def getNewsFeed(self, userId: int) -> list[int]:
        res = []
        max_heap = []

        self.follow_list[userId].add(userId)
        for followeeId in self.follow_list[userId]:
            if followeeId in self.tweets:
                i = len(self.tweets[followeeId]) - 1
                count, tweetId = self.tweets[followeeId][i]
                heapq.heappush_max(max_heap, [count, tweetId, followeeId, i - 1])

        while max_heap and len(res) < 10:
            count, tweetId, followeeId, i = heapq.heappop_max(max_heap)
            res.append(tweetId)

            if i >= 0:
                count, tweetId = self.tweets[followeeId][i]
                heapq.heappush_max(max_heap, [count, tweetId, followeeId, i - 1])

        return res

    def follow(self, followerId: int, followeeId: int) -> None:
        self.follow_list[followerId].add(followeeId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        self.follow_list[followerId].discard(followeeId)
