from collections import Counter


class ValidAnagram:
    def isAnagram(self, s: str, t: str) -> bool:
        return Counter(s) == Counter(t)

# 1. carification
# 2. possible solutions ->optimal (time & space)
# 3  code
# test case
