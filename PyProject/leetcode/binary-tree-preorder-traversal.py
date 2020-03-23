from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class BinaryTreePreorderTraversal:
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        res = []
        if root:
            res.append(root.val)
            res.extend(self.preorderTraversal(root.left))
            res.extend(self.preorderTraversal(root.right))
        return res
