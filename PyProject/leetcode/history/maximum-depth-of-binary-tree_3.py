from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(root: TreeNode) -> int:
    if root is None:
        return 0
    else:
        left_dep = helper(root.left)
        right_dep = helper(root.right)
        return max(left_dep, right_dep) + 1


# 自底向上
class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        return helper(root)
