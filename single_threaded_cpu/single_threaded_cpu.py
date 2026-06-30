import heapq


class Solution:
    def getOrder(self, tasks: list[list[int]]) -> list[int]:
        res = []

        enqueue_time_heap = []
        for i, (enqueue_time, process_time) in enumerate(tasks):
            enqueue_time_heap.append([enqueue_time, process_time, i])

        heapq.heapify(enqueue_time_heap)
        processing_time_heap = []

        i = enqueue_time_heap[0][0]
        while enqueue_time_heap or processing_time_heap:
            while enqueue_time_heap and enqueue_time_heap[0][0] <= i:
                _, p_time, idx = heapq.heappop(enqueue_time_heap)
                heapq.heappush(processing_time_heap, [p_time, idx])

            if processing_time_heap:
                p_time, idx = heapq.heappop(processing_time_heap)
                res.append(idx)
                i += p_time

            else:
                if len(enqueue_time_heap) != 0:
                    i = enqueue_time_heap[0][0]

        return res
