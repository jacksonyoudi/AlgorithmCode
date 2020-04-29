from typing import List


def partition(nums: List[int], low: int, high: int):
    i = low - 1
    pivot = nums[high]

    for j in range(low, high):
        if nums[j] > pivot:
            i += 1
            nums[i], nums[j] = nums[j], nums[i]
    nums[i + 1], nums[high] = nums[high], nums[i + 1]
    return i + 1


def quick_sort(nums: List[int], low: int, high: int):
    if high < low:
        pi = partition(nums, low, high)
        quick_sort(nums, low, pi - 1)
        quick_sort(nums, pi + 1, high)


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if len(nums) == 1:
            return nums

        if nums[0] > nums[1]:
            nums.reverse()
            return nums

        l = len(nums) - 2

        for i in range(l):
            pass
