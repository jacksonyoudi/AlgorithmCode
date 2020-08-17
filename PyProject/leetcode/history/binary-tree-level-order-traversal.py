from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(nodes: List[TreeNode], result: List[List[int]]):
    next_nodes = []
    cur_vals = []
    for node in nodes:
        if node is None:
            continue
        cur_vals.append(node.val)
        next_nodes.append(node.left)
        next_nodes.append(node.right)
    if cur_vals:
        result.append(cur_vals)
        helper(next_nodes, result)


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        result = []
        nodes = [root]
        helper(nodes, result)
        return result
