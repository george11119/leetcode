import heapq


class Solution:
    def findKthLargest(self, nums: list[int], k: int) -> int:
        heapq.heapify_max(nums)

        for i in range(k):
            res = heapq.heappop_max(nums)
            if i == k - 1:
                return res

        return -1
