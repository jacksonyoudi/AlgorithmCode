# -*- coding:utf-8 -*-
# author: youdi

from typing import List


def bubble_sort(nums: List[int]) -> List[int]:
    nums_length = len(nums)
    for i in range(1, nums_length):
        for j in range(0, nums_length - i):
            if nums[j] > nums[j + 1]:
                nums[j], nums[j + 1] = nums[j + 1], nums[j]
    return nums


if __name__ == '__main__':
    A = [64, 25, 12, 22, 11]
    B = bubble_sort(A)
    print(B)
