from typing import List


class Subsets:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []
        n = len(nums)

        def helper(i, tmp):
            res.append(tmp)
            for j in range(i, n):
                helper(j + 1, tmp + [nums[j]])

        helper(0, [])
        return res
