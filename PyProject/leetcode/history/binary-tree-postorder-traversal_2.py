from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        result = []
        if root:
            result.extend(self.postorderTraversal(root.left))
            result.extend(self.postorderTraversal(root.right))
            result.append(root.val)

        return result
