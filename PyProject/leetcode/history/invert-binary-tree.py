class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class InvertBinaryTree:
    def invertTree(self, root: TreeNode) -> TreeNode:
        if root:
            self.invertTree(root.left)
            self.invertTree(root.right)
            root.left, root.right = root.right, root.left
        return root
