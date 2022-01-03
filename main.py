import data_structure as ds

ll = ds.MyLinkedList()
ll.addAtHead(1)
ll.addAtTail(3)
ll.addAtIndex(1, 2)
tmp = ll.get(1)
print(tmp)

ll.deleteAtIndex(1)
tmp = ll.get(1)

print(tmp)
# ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
# [[],[1],[3],[1,2],[1],[1],[1]]
