# coding=utf-8


"""
# Employee info
class Employee:
    def __init__(self, id, importance, subordinates):
        # It's the unique id of each node.
        # unique id of this employee
        self.id = id
        # the importance value of this employee
        self.importance = importance
        # the id of direct subordinates
        self.subordinates = subordinates
"""


class Solution:
    def getImportance(self, employees, id):
        """
        :type employees: Employee
        :type id: int
        :rtype: int
        """
        dicts = {}
        for employee in employees:
            dicts[employee.id] = employee
        score = 0
        subs = [id]
        while subs:
            employee = dicts[subs[0]]
            score = score + employee.importance
            del subs[0]
            for subid in employee.subordinates:
                subs.append(subid)
        return score
