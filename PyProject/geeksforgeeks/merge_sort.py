# -*- coding:utf-8 -*-
# author: youdi

from typing import List


def merge_sort(nums: List[int]) -> List[int]:
    if len(nums) > 1:
        m = len(nums) // 2
        left = nums[:m]
        right = nums[m:]

        left = merge_sort(left)
        right = merge_sort(right)

        nums = []

        while len(left) > 0 and len(right) > 0:
            if left[0] < right[0]:
                nums.append(left[0])
                left.pop(0)
            else:
                nums.append(right[0])
                right.pop(0)
        if len(left) > 0:
            nums.extend(left)
        if len(right) > 0:
            nums.extend(right)
    return nums


if __name__ == '__main__':
    A = [64, 25, 12, 22, 11]
    B = merge_sort(A)
    print(B)
