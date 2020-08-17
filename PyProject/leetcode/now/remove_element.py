# coding: utf-8

from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        # 双指针法
        one = 0
        two = 1

        for i in nums:
            if i == val:
                nums[one] = two
