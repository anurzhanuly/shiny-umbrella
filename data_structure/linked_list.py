from typing import Type

from leetcode.data_structure.list_node import ListNode

class LinkedList:
    head: ListNode = None

    def insert(self, head: ListNode, node: ListNode):
        if self.head:
            node.next_node = self.head
        
        self.head = node

    def delete(self, head: ListNode, node: ListNode):
        pass
        
            



    