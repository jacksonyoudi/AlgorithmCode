import math

class Solution:
    def subarrayGCD(self, nums: List[int], k: int) -> int:
        cnt = 0
        for i in range(len(nums)):
            for j in range(i, len(nums)):
                if self.findGCD(nums[i:j+1]) == k:
                    cnt = cnt + 1
        return cnt


    def findGCD(self, nums: List[int]) -> int:
        mx, mn = max(nums), min(nums)
        return math(mx, mn)
