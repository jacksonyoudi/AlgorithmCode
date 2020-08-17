class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class LowestCommonAncestorOfABinaryTree:
    def __init__(self):
        self.ans = None

    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        # 闭包
        def recurse_tree(node: TreeNode) -> bool:
            if not node:
                return False
            left = recurse_tree(node.left)
            right = recurse_tree(node.right)

            mid = node == q or node == p

            if mid + left + right >= 2:
                self.ans = node
            return mid or left or right

        recurse_tree(root)
        return self.ans
