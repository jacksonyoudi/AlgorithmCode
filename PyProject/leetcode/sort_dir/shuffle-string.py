from typing import List


class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        result = ['' for i in range(len(s))]
        for i, v in enumerate(indices):
            result[v] = s[i]
        return ''.join(result)
