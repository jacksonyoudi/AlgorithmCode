# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/6/3 23:18
# software: IntelliJ IDEA

from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(root: TreeNode, result: List[int]):
    if root is None:
        return
    else:
        helper(root.left, result)
        helper(root.right, result)
        result.append(root.val)


class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        result = []
        helper(root, result)
        return result
