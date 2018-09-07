# coding=utf-8


"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""


class Solution(object):
    def maxDepth(self, root):
        """
        :type root: Node
        :rtype: int
        """
        if not root:
            return 0

        depth = 1
        childdepth = []
        for child in root.children:
            childdepth.append(self.maxDepth(child))

        if childdepth:
            if len(childdepth) > 1:
                return depth + max(*childdepth)
            else:
                return depth + childdepth[0]
        return depth
