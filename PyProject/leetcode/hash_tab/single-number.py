from typing import List
from functools import reduce


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        # 函数式编程
        return reduce(lambda x, y: x ^ y, nums)
