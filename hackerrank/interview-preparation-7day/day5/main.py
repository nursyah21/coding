class SinglyLinkedListNode:
    def __init__(self, node_data):
        self.data = node_data
        self.next = None

class SinglyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None

    def insert_node(self, node_data):
        node = SinglyLinkedListNode(node_data)
        if not self.head:
            self.head = node
        else:
            self.tail.next = node


        self.tail = node
    
    def printlist(self):
        temp = self.head
        while temp:
            print(temp.data, end=" ")
            temp = temp.next


## non recursive
def mergeLists(headA, headB):
    #dummy node to store the result
    result = SinglyLinkedListNode(None)
    tail = result

    while headA and headB:
        if headA.data < headB.data:
            tail.next = headA
            headA = headA.next
        else:
            tail.next = headB
            headB = headB.next
        tail = tail.next
    
    if headA:
        tail.next = headA
    elif headB:
        tail.next = headB
    
    return result.next

## runtime error
def mergeLists(headA, headB):
    #dummy node to store the result
    if headA is None:
        return headB
    if headB is None:
        return headA
    
    if headA.data < headB.data :
            headA.next = mergeLists(headA.next, headB)
            return headA
    else:
            headB.next = mergeLists(headA, headB.next)
            return headB


def printlist(head):
    while head:
        print(head.data, end=" ")
        head = head.next

def isBalanced(s):
  if len(s) % 2 != 0:
    return 'NO'
  
  stack = []
  closing = ['}', ']', ')']
  pairs = {'}':'{', ']':'[', ')':'('}
  
  for bracket in s:
    if bracket in closing:
      if not stack or stack.pop() != pairs.get(bracket):
        return 'NO'
    else:
      stack.append(bracket)
      
  return 'YES' if not stack else 'NO'
    

if __name__ == "__main__":
    print(isBalanced("{{[]}[]()}"))

    """ headA = SinglyLinkedList()
    headB = SinglyLinkedList()
    headA.insert_node(10)
    #headA.insert_node(15)
    headB.insert_node(5)
    #headB.insert_node(50)
    headC = SinglyLinkedList()
    """
    
    # res = mergeList(headA, headC)
    # res.printlist()
    # print()

    """ headA.printlist()
    print()
    headB.printlist()
    print()
    res = mergeLists(headA.head, headB.head)

    printlist(res) """
    #res.printlist()
    # print()
    