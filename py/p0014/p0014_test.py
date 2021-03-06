# coding=utf-8

import unittest
from .p0014 import Solution


class TestP0014(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        pass

    @classmethod
    def tearDownClass(cls):
        pass

    def setUp(self):
        self.solution = Solution()

    def tearDown(self):
        self.solution = None

    def test_p0014(self):
        self.assertEqual(
            "fl",
            self.solution.longestCommonPrefix(["flower", "flow", "flight"])
        )
