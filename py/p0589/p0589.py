# coding=utf-8


"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""


class Solution(object):
    def preorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        if not root:
            return []

        nodes = [root]
        ret = list()
        while nodes:
            node = nodes[-1]
            del nodes[-1]
            ret.append(node.val)
            for i in range(len(node.children)):
                if node.children[len(node.children)-i-1]:
                    nodes.append(node.children[len(node.children)-i-1])

        return ret
