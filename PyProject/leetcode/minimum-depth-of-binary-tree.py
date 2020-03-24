class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class MaximumDepthOfBinaryTree:
    def minDepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        if not root.left:
            return self.minDepth(root.right) + 1
        if not root.right:
            return self.minDepth(root.left) + 1
        return min(self.minDepth(root.right), self.minDepth(root.left)) + 1
