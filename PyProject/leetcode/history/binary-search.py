from typing import List


def helper(nums: List[int], low: int, high: int, target: int) -> int:
    if low > high:
        return -1
    mid = low + (high - low >> 1)
    if nums[mid] > target:
        return helper(nums, low, mid - 1, target)
    elif nums[mid] < target:
        return helper(nums, mid + 1, high, target)
    else:
        return mid


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return helper(nums, 0, len(nums) - 1, target)
