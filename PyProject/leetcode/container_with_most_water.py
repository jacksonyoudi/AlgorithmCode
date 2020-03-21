from typing import List


class ContainerWithMostWater:
    def maxArea(self, height: List[int]) -> int:
        max = 0
        start = 0
        end = len(height) - 1

        while start < end:
            if height[start] < height[end]:
                max_tmp = height[start] * (end - start)
                start += 1
            else:
                max_tmp = height[end] * (end - start)
                end -= 1
            if max_tmp > max:
                max = max_tmp
        return max

    def maxArea2(self, height: List[int]) -> int:
        start, end, area = 0, len(height) - 1, 0
        while start < end:
            if height[start] < height[end]:
                res = max(height[start] * (end - start), res)
                start += 1
            else:
                res = max(height[end] * (end - start), res)
                end -= 1
        return res
