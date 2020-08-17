from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(root: TreeNode, dep: int):
    if root:
        left_dep = helper(root.left, dep + 1)
        right_dep = helper(root.right, dep + 1)
        dep = max(left_dep, right_dep)
    return dep


# 自顶向下
class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        return helper(root, 0)
