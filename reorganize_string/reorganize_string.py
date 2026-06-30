import heapq
from collections import defaultdict, Counter


class Solution:
    def reorganizeString(self, s: str) -> str:
        freq_map = defaultdict(int)

        for c in s:
            freq_map[c] += 1

        max_heap = []
        for c, freq in freq_map.items():
            max_heap.append([freq, c])

        heapq.heapify_max(max_heap)

        prev = None
        res = ""
        while max_heap or prev:
            if prev and not max_heap:
                return ""

            freq, c = heapq.heappop_max(max_heap)
            res += c
            freq -= 1

            if prev:
                heapq.heappush_max(max_heap, prev)
                prev = None

            if freq != 0:
                prev = [freq, c]

        return res
