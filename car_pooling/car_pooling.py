import heapq


class Solution:
    def carPooling(self, trips: list[list[int]], capacity: int) -> bool:
        trips.sort(key=lambda t: t[1])

        min_heap = []  # [end pass, num passengers]
        cur_pass = 0

        for num_pass, start, end in trips:
            while min_heap and min_heap[0][0] <= start:
                cur_pass -= heapq.heappop(min_heap)[1]

            cur_pass += num_pass
            if cur_pass > capacity:
                return False

            heapq.heappush(min_heap, [end, num_pass])

        return True
