from typing import List


class Combinations:
    def combine(self, n: int, k: int) -> List[List[int]]:
        return [[i, j] for j in range(1, n + 1) for i in range(1, n + 1) if i < j]
