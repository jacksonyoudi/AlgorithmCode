class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class ValidateBinarySearchTree:
    def isValidBST(self, root: TreeNode) -> bool:
        return self.helper(root)

    def helper(self, root: TreeNode, min=float('-inf'), max=float('inf')) -> bool:
        if not root:
            return True
        if not min < root.val < max:
            return False
        return self.helper(root.left, min, root.val) and self.helper(root.right, root.val, max)
