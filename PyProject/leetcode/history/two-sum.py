from typing import List


class TwoSum:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        m = dict()
        for i, v in enumerate(nums):
            complementary = target - v
            if m.get(complementary) is not None:
                return [m.get(complementary), i]
            m[v] = i
