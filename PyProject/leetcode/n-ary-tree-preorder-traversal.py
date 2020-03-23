from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class NAryTreePreorderTraversal:
    def preorder(self, root: 'Node') -> List[int]:
        res = []
        if root:
            res.append(root.val)
            for i in root.children:
                res.extend(self.preorder(i))
        return res
