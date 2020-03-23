from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class BinaryTreeInorderTraversal:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        res = []
        self.helper(root, res)
        return res

    def helper(self, root: TreeNode, res: List[int]):
        if root:
            self.helper(root.left, res)
            res.append(root.val)
            self.helper(root.right, res)
