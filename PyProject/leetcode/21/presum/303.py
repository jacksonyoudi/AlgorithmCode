from typing import List
from functools import reduce


class NumArray:
    def __init__(self, nums: List[int]):
        self.nums = nums

    def sumRange(self, left: int, right: int) -> int:
        return sum(self.nums[left:right + 1])


if __name__ == '__main__':
    a = [1, 2, 3, 4, 5]
    b = reduce(lambda x, y: x + y, a)
    print(b)
