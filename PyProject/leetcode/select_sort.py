from typing import List


class SelectSort:
    def select_sort(l: List[int]) -> List[int]:
        length = len(l)
        max_val = -10
        for i in range(length):
            for j in range(length - i):
                pass


class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        # 选择排序， 每次找到最大的，放在对应位置
        length = len(nums)
        for i in range(length):
            max_index = i
            for j in range(0, length - i):
                if nums[j] > nums[max_index]:
                    max_index = j
            if max_index != i:
                nums[max_index], nums[length - i] = nums[length - i], nums[max_index]
        return nums
