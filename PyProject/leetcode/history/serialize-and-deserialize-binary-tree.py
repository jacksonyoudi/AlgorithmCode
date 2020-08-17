from typing import List


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Codec:
    def serialize(self, root: TreeNode):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """

        def tostr(s: str, root: TreeNode) -> str:
            if not root:
                s += 'None,'
                return s
            s += str(root.val) + ','
            s = tostr(s, root.left)
            s = tostr(s, root.right)
            return s

        return tostr("", root)

    def deserialize(self, data: str) -> TreeNode:
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """

        def to_node(q: List[str]) -> TreeNode:
            if l[0] == 'None':
                l.pop(0)
                return None
            v = q.pop(0)
            root = TreeNode(int(v))
            root.left = to_node(l)
            root.right = to_node(l)
            return root

        l = data.split(",")
        l = [i for i in l if i]
        return to_node(l)
