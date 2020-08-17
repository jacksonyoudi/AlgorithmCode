class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class MaximumDepthOfBinaryTree:
    def maxDepth(self, root: TreeNode) -> int:
        return self.helper(root, 0)

    def helper(self, root: TreeNode, level: int) -> int:
        if not root:
            return level
        left_level = self.helper(root.left, level + 1)
        right_level = self.helper(root.right, level + 1)
        if left_level > right_level:
            return left_level
        else:
            return right_level
