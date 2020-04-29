from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        start, end, area = 0, len(height) - 1, 0
        while start < end:
            if height[start] > height[end]:
                area = max(height[end] * (end - start), area)
                end -= 1
            else:
                area = max(height[start] * (end - start), area)
                start += 1
        return area
