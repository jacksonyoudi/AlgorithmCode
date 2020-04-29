from typing import List


def partition(nums: List[int], low: int, high: int) -> int:
    i = low - 1
    pivot = nums[high]
    for j in range(low, high):
        if nums[j] < pivot:
            i += 1
            nums[i], nums[j] = nums[j], nums[i]
    nums[i + 1], nums[high] = nums[high], nums[i + 1]
    return i + 1


def quick_sort(nums: List[int], low: int, high: int):
    if low < high:
        pivot = partition(nums, low, high)
        quick_sort(nums, low, pivot - 1)
        quick_sort(nums, pivot + 1, high)


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        result = []
        length = len(nums)
        if not nums or length < 3:
            return result

        # 先排序
        quick_sort(nums, 0, len(nums) - 1)

        for i, num in enumerate(nums):
            if num > 0:
                return result
            # 去掉重复数据
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            start = i + 1
            end = length - 1
            while start < end:
                if nums[i] + nums[start] + nums[end] == 0:
                    result.append([nums[i], nums[start], nums[end]])

                    while start < end and nums[start] == nums[start + 1]:
                        start += 1
                    while start < end and nums[end] == nums[end - 1]:
                        end -= 1
                    start += 1
                    end -= 1
                elif nums[i] + nums[start] + nums[end] > 0:
                    end -= 1
                else:
                    start += 1
        return result
