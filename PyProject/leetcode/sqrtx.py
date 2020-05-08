class Solution:
    def mySqrt(self, x):
        if x < 2:
            return x

        left = self.mySqrt(x >> 2) << 1
        right = left + 1
        return left if right * right > x else right
