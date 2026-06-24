import heapq


class Solution:
    def lastStoneWeight(self, stones: list[int]) -> int:
        heapq.heapify_max(stones)

        while len(stones) > 1:
            larger_stone = heapq.heappop_max(stones)
            smaller_stone = heapq.heappop_max(stones)

            if larger_stone == smaller_stone:
                continue
            else:
                larger_stone -= smaller_stone
                heapq.heappush_max(stones, larger_stone)

        return 0 if len(stones) == 0 else stones[0]
