#!/usr/bin/env python3

from typing import List
from .list_node import ListNode

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


s = Stack()
n = ListNode()
n.value = 12
print(n)
s.push()
