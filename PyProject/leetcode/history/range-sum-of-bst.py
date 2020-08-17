# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(root: TreeNode, L: int, R: int, result: int) -> int:
    if root is None:
        return result
    if L <= root.val <= R:
        result += root.val
    result = helper(root.left, L, R, result)
    result = helper(root.right, L, R, result)
    return result


class Solution:
    def rangeSumBST(self, root: TreeNode, L: int, R: int) -> int:
        result = helper(root, L, R, 0)
        return result
