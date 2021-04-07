func sortedInsert(head *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
    newNode := DoublyLinkedListNode{data: data}
    if (head == nil) {
        return &newNode
    }
    cur := head
    var prev *DoublyLinkedListNode = nil
    
    if(data < cur.data){
        newNode.next = cur
        newNode.prev = nil
        cur.prev = &newNode
        return &newNode
    }
    
    for (cur != nil && data > cur.data) {
        prev = cur
        cur = cur.next
    }
    
    if (cur == nil) {
        newNode.prev = prev
        newNode.next = nil
        prev.next = &newNode
        return head
    } else {
        newNode.prev = prev
        newNode.next = prev.next
        prev.next = &newNode
        newNode.next.prev = &newNode
    }
    return head
}