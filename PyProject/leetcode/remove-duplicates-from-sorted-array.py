from typing import List


class RemoveDuplicatesFromSortedArray:
    def removeDuplicates(self, nums: List[int]) -> int:
        start = 0
        end = 1
        length = len(nums)
        while end < length - 1:
            if (end + 1) >= length:
                break
            if nums[end] == nums[end + 1]:
                end += 1
                continue
            else:
                nums[start + 1] = nums[end + 1]
                start += 1
                end += 1
        return start
