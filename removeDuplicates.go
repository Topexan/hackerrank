func removeDuplicates(head *SinglyLinkedListNode) *SinglyLinkedListNode {
    cur := head
    for (cur != nil && cur.next != nil ) {
        if (cur.next.data == cur.data) {
            cur.next = cur.next.next
        } else {
            cur = cur.next
        }
    }
    return head
}