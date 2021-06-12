from typing import List


class Solution:
    def chalkReplacer(self, chalk: List[int], k: int) -> int:
        if len(chalk) == 1:
            return 1

        s = sum(chalk)
        if s < k:
            k = k % s

        while k >= 0:
            for i, v in enumerate(chalk):
                k = k - v
                if k < 0:
                    return i

        return 0
