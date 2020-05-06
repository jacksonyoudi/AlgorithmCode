from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(root: TreeNode, result: List[int]) -> List[int]:
    if root is None:
        return result
    else:
        result.append(root.val)
        result = helper(root.right, result)
        result = helper(root.left, result)
    return result


class Solution:
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        result = []
        result = helper(root, result)
        return result
