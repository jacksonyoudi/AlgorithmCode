from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class NaryTreePostorderTraversal:
    def postorder(self, root: Node) -> List[int]:
        res = []
        if root:
            for i in self.children:
                res.extend(self.postorder(i))
            res.append(root.val)
        return res
