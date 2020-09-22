# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def __init__(self):
        self.sum = 0

    def convertBST(self, root: TreeNode) -> TreeNode:
        if not root:
            return root

        self.convertBST(root.right)

        self.sum += root.val
        root.val = self.sum

        self.convertBST(root.left)
        return root
