# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def helper(a: int, root: TreeNode) -> int:
    if not root:
        return a
    else:
        a += root.val
        b = helper(a, root.left)
        c = helper(b, root.right)
    return c


class Solution:
    def __init__(self):
        self.__sum = 0

    def findTilt(self, root: TreeNode) -> int:
        if root.left is None and root.right is None:
            return
        elif root.left is not None and root.right is None:
            self.__sum += abs(helper(0, root.left))
            self.findTilt(root.left)
        elif root.left is None and root.right is not None:
            self.__sum += abs(helper(0, root.right))
            self.findTilt(root.right)
        else:
            self.__sum += abs(helper(0, root.left) -helper(0, root.right))
            self.findTilt(root.right)
            self.findTilt(root.left)
        return self.__sum
