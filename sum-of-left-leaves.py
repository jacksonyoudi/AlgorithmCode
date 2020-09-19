# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(root: TreeNode, result: int, flags: bool):
    if not root:
        return result
    if flags and not root.right and not root.left:
        result += root.val
    result = helper(root.left, result, True)
    result = helper(root.right, result, False)
    return result


class Solution:
    def sumOfLeftLeaves(self, root: TreeNode) -> int:
        return helper(root, 0, False)
