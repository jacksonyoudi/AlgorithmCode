from typing import List

"""
使用排序 + 加双指针法

"""


class ThresSum:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # 特判
        res = []
        length = len(nums)
        if not nums or length < 3:
            return res

        # 排序
        nums.sort()

        for i, num in enumerate(nums):
            # 如果num > 0 说明后面都是大于0 剪枝
            if num > 0:
                return res
            #  重复的数据
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            start = i + 1
            end = length - 1
            while start < end:
                if nums[i] + nums[start] + nums[end] == 0:
                    res.append([nums[i], nums[start], nums[end]])
                    # 如果出现重复数据
                    while start < end and nums[start] == nums[start + 1]:
                        start += 1
                    while start < end and nums[end] == nums[end - 1]:
                        end -= 1
                    start += 1
                    end -= 1
                elif nums[i] + nums[start] + nums[end] > 0:
                    end -= 1
                else:
                    start += 1
        return res
