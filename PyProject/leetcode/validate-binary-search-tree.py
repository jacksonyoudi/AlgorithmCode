class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class ValidateBinarySearchTree:
    def isValidBST(self, root: TreeNode) -> bool:
        # terminator
        if root:
            if root.left:
                if root.val < root.left.val:
                    return False

            if root.right:
                if root.val > root.right.val:
                    return False

            if root.left and root.right:
                if root.left.val > root.right.val:
                    return False

            if not self.isValidBST(root.left):
                return False

            if not self.isValidBST(root.right):
                return False

        return True
