from typing import List


class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        pairs = 0

        if not nums:
            return pairs
        l = len(nums)
        for i, v in enumerate(nums):
            for j in nums[i:l]:
                if nums[i] == j:
                    pairs += 1
        return pairs
