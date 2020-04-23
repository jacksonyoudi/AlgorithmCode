from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(heads: List[TreeNode], result: List[int]):
    tmp = []
    nodes = []
    for i in heads:
        if i is not None:
            tmp.append(i.val)
        else:
            continue
        if i.left is not None:
            nodes.append(i.left)
        if i.right is not None:
            nodes.append(i.right)
    if len(tmp) == 0:
        return
    else:
        result.append(tmp[-1])
        helper(nodes, result)


class Solution:
    def rightSideView(self, root: TreeNode) -> List[int]:
        result = []
        helper([root], result)
        return result
