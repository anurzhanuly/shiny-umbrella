class LinkedList:

    def __init__(self):
        self.head = None
        self.tail = None
        self.index = -1

    def get(self, index: int) -> int:
        if index > self.index or not self.head or index < 0:
            return -1

        node = self.head

        while index:
            node = node.next
            index -= 1

        return node.val

    def addAtHead(self, val: int) -> None:
        node = Node(val)

        if not self.head:
            self.head = node
            self.tail = self.head
        else:
            node.next = self.head
            self.head = node

        self.index += 1

    def addAtTail(self, val: int) -> None:
        node = Node(val)

        if not self.tail:
            self.head = node
            self.tail = self.head
        else:
            self.tail.next = node
            self.tail = node

        self.index += 1

    def addAtIndex(self, index: int, val: int) -> None:
        if index == 0:
            return self.addAtHead(val)

        if index > self.index or index < 0:
            return

        if index == self.index:
            return self.addAtTail(val)

        node = self.head

        while index != 1:
            index -= 1
            node = node.next

        tmp = node.next
        node.next = Node(val)
        node.next.next = tmp

        self.index += 1

    def deleteAtIndex(self, index: int) -> None:
        if index > self.index or not self.head or index < 0:
            return

        node = self.head

        while index != 1:
            node = node.next
            index -= 1

        node.next = node.next.next
        self.index -= 1


class Node:

    def __init__(self, val: int = 0):
        self.val = val
        self.next = None