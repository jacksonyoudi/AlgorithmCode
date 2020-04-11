from typing import List


class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        if n <= 1:
            return nums

        # 默认使用第一个数当已经排序好的数
        for i in range(1, n):
            j = i
            target = nums[i]
            while j > 0 and target < nums[j - 1]:
                nums[j] = nums[j - 1]
                j = j - 1
        return nums
