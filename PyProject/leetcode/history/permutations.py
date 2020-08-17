# coding: utf8
from copy import deepcopy


def helper(pre, nums, result):
    if len(pre) == len(nums):
        result.append(pre)
    for i in set(nums).difference(set(pre)):
        i_tmp = deepcopy(pre)
        i_tmp.append(i)
        helper(i_tmp, nums, result)


if __name__ == '__main__':
    a = [1, 2, 3]
    result = []
    helper([], a, result)
    print(result)