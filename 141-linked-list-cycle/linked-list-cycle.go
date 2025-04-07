/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next.Next == head {
		return true
	}

	walker := head
	runner := head
	for walker != nil && runner != nil {
		walker = walker.Next
		if runner.Next != nil {
			runner = runner.Next.Next
		}else {
            return false
        }
		if runner == walker {
			return true
		}
	}

	return false

}