from typing import List
from functools import reduce


class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        # result = []
        # if not nums:
        #     return result
        # for i in nums:
        #     if result:
        #         result.append(result[-1] + i)
        #     else:
        #         result.append(i)
        #
        # return result

        return [sum(nums[:i + 1]) for i in range(len(nums))]


def add(a: int, b: int):
    return a + b


if __name__ == '__main__':
    a = [1, 2, 3, 4]
    c = reduce(add, a)
    print(c)
