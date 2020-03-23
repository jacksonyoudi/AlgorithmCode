from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class NAryTreeLevelOrderTraversal:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        q, ret = [root], []
        while any(q):
            ret.append([node.val for node in q])
            q = [child for node in q for child in node.children if child]
        return ret

    def levelOrder2(self, root: 'Node') -> List[List[int]]:
        # 当做队列
        q = [root]
        res = []
        while q:
            temp = []
            next_stack = []
            for node in q:
                temp.append(node.val)
                for child in node.children:
                    next_stack.append(child)
            res.append(temp)
            q = next_stack
        return res
