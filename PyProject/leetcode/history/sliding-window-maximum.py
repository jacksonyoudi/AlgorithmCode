from typing import List


class SlidingWindowMaximum:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        wind_max = []
        for i in range(len(nums) - k + 1):
            wind_max.append(max(nums[i: i + k]))
        return wind_max
