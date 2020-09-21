# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def dfs(root: TreeNode):
    nonlocal total
    if root:
        dfs(root.right)
        total += root.val
        root.val = total
        dfs(root.left)


class Solution:
    def convertBST(self, root: TreeNode) -> TreeNode:
        total = 0
        dfs(root)
        return root
