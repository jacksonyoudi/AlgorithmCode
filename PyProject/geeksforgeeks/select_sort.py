# -*- coding:utf-8 -*-
# author: youdi

from typing import List


def select_sort(nums: List[int]) -> List[int]:
    nums_length = len(nums)
    for i in range(nums_length):
        min_index = i
        for j in range(i + 1, nums_length):
            if nums[j] < nums[min_index]:
                min_index = j
        nums[min_index], nums[i] = nums[i], nums[min_index]
    return nums


if __name__ == '__main__':
    A = [64, 25, 12, 22, 11]
    B = select_sort(A)
    print(B)
