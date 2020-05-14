from typing import List


def main(nums: List[int]):
    # nums.reverse()
    s = slice(len(nums), 0, -1)
    nums = nums[s]
    print(nums)


if __name__ == '__main__':
    a = [1, 2, 3, 4, 5]
    main(a)
