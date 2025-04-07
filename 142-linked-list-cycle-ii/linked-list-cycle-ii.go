/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == head {
		return head
	}

	walker := head
	runner := head
	for walker != nil && runner != nil {
		walker = walker.Next
		if runner.Next != nil {
			runner = runner.Next.Next
		} else {
			break
		}
		if runner == walker {
			break
		}
	}

	if runner == nil || runner.Next == nil {
		return nil
	}

	newS := head

	for newS != walker {
		newS = newS.Next
		walker = walker.Next
	}

	return newS

}

