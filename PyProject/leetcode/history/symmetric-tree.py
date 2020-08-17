from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def is_symmetric(left: TreeNode, right: TreeNode) -> bool:
    if left is None and right is None:
        return True
    if left is None or right is None:
        return False

    if left.val == right.val:
        out = is_symmetric(left.left, right.right)
        inner = is_symmetric(left.right, right.left)
        return out and inner
    else:
        return False


class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        if root is None:
            return True
        else:
            return is_symmetric(root.left, root.right)
