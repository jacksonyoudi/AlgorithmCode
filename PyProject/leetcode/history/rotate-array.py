from typing import List


class RotateArray:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        self.reversed(nums, 0, len(nums) - 1)
        self.reversed(nums, 0, k - 1)
        self.reversed(nums, k, len(nums) - 1)

    def reversed(self, nums: List[int], start, end):
        while start < end:
            # 交换
            tmp = nums[start]
            nums[start] = nums[end]
            nums[end] = tmp
            start += 1
            end -= 1
