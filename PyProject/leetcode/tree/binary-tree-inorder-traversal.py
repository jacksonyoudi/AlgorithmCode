from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(node: TreeNode, result: List[int]) -> List[int]:
    if not node:
        return result
    helper(node.left, result)
    result.append(node.val)
    helper(node.right, result)
    return result


class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        result = []
        return helper(root, result)
