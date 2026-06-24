import heapq


class Solution:
    def leastInterval(self, tasks: list[str], n: int) -> int:
        tasks_count = {}
        for t in tasks:
            tasks_count[t] = tasks_count.get(t, 0) + 1

        max_heap = [v for _, v in tasks_count.items()]

        heapq.heapify_max(max_heap)
        q = []

        i = 0
        while len(max_heap) != 0 or len(q) != 0:
            i += 1

            if len(max_heap) != 0:
                cnt = heapq.heappop_max(max_heap) - 1
                if cnt != 0:
                    q.append([cnt, i + n])

            if len(q) != 0 and q[0][1] == i:
                heapq.heappush_max(max_heap, q.pop(0)[0])

        return i
