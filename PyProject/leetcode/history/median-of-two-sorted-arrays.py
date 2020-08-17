from typing import List


# merge
def merge(nums: List[int]):
    if len(nums) > 1:
        mid = len(nums) // 2
        L = nums[:mid]
        R = nums[mid:]
        merge(L)
        merge(R)
        i = j = k = 0
        while i < len(L) and j < len(R):
            if L[i] > R[j]:
                nums[k] = R[j]
                j += 1
            else:
                nums[k] = L[i]
                i += 1
            k += 1

        while i < len(L):
            nums[k] = L[i]
            i += 1
            k += 1

        while j < len(R):
            nums[k] = R[j]
            j += 1
            k += 1


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        nums1.extend(nums2)
        merge(nums1)
        l = len(nums1)
        print(l // 2)
        print(nums1)
        if l % 2 == 0:
            return (nums1[l // 2 - 1] + nums1[l // 2]) / 2
        else:
            return nums1[l // 2]


if __name__ == '__main__':
    a = [1, 2]
    b = [3, 4]
    c = Solution()
    c.findMedianSortedArrays(a, b)
