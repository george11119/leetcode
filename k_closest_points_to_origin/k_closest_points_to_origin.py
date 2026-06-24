import heapq


# class Point(list[int]):
#     def __lt__(self, other):
#         return ((self[0]) ** 2 + (self[1]) ** 2) < ((other[0]) ** 2 + (other[1]) ** 2)
#
#
# class Solution:
#     def kClosest(self, points: list[list[int]], k: int) -> list[list[int]]:
#         heap = [Point(p) for p in points]
#         heapq.heapify(heap)
#
#         res = []
#         for i in range(k):
#             res.append(heapq.heappop(heap))
#
#         return res

# turns out you can just have the compare value in the first part of the list
# instead of writing a custom compare method
class Solution:
    def kClosest(self, points: list[list[int]], k: int) -> list[list[int]]:
        min_heap = [[x ** 2 + y ** 2, x, y] for [x, y] in points]
        heapq.heapify(min_heap)

        res = []
        for i in range(k):
            _, x, y = heapq.heappop(min_heap)
            res.append([x, y])

        return res
