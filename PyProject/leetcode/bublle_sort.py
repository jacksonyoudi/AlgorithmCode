from typing import List


class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        length = len(nums)
        # 优化，如果一次都没有交换

        # 冒泡排序特点
        for i in range(length):
            is_swap = False
            for j in range(length - i - 1):
                # 判断大小并交换
                if nums[j] > nums[j + 1]:
                    is_swap = True
                    nums[j], nums[j + 1] = nums[j + 1], nums[j]
            if is_swap is False:
                return nums
            return nums
