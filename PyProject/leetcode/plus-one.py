from typing import List


class PlusOne:
    def addOne(self, digits: List[int], index: int) -> List[int]:
        s = digits[index]
        if s < 10:
            digits[index] = s
            return digits
        else:
            digits[index] = 0
            if index == 0:
                digits.insert(0, 0)
                index = 1
            return self.addOne(digits, index - 1)

    def plusOne(self, digits: List[int]) -> List[int]:
        return self.addOne(digits, len(digits) - 1)
