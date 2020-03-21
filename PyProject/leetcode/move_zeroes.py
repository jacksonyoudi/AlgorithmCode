from typing import List


class MoveZeroes:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        使用双指针法
        j 记录出现0的位置
        i 记录非0的位置
        """
        j = 0
        for i, v in enumerate(nums):
            # 如果出现第0个，那就自身换自身
            if v != 0:
                nums[j], nums[i] = nums[i], nums[j]
                j += 1
