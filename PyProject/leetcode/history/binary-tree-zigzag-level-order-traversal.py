# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/6/3 23:28
# software: IntelliJ IDEA


from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def helper(nodes: List[TreeNode], result: List[List[int]], reverse: bool):
    tmp = []
    new_nodes = []
    for i in nodes:
        tmp.append(i.val)
        if i.left:
            new_nodes.append(i.left)
        if i.right:
            new_nodes.append(i.right)
    if len(tmp) != 0:
        if reverse:
            result.append(tmp[::-1])
        else:
            result.append(tmp)
    if len(new_nodes) != 0:
        helper(new_nodes, result, not reverse)


class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        nodes = [root]
        result = []
        helper(nodes, result, False)
        return result
