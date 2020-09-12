from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(nodes: List[TreeNode], result: List[float]) -> List[float]:
    if not nodes:
        return result
    l = []
    r = []
    for node in nodes:
        r.append(node.val)
        if node.left:
            l.append(node.left)
        if node.right:
            l.append(node.right)
    result.append(
        sum(r) / len(r)
    )
    return helper(l, result)


class Solution:
    def averageOfLevels(self, root: TreeNode) -> List[float]:
        if not root:
            return 0
        return helper([root], [])
