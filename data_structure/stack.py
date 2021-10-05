from typing import List
from data_structure.list_node import ListNode

class Stack:
    top: ListNode = None

    def push(self, node: ListNode):
        if self.top:
            node.next = self.top
        else:
            self.top = node

    def pop(self):
        if self.top:
            tmp = self.top
            self.top = self.top.next

            return tmp
        else:
            return None
