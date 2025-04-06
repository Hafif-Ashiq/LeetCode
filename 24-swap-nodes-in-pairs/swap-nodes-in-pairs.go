/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	iter := head

	for iter != nil {

		if iter.Next != nil {

			if iter == head {
				iterNext := iter.Next
				iter.Next = iterNext.Next
				iterNext.Next = iter
				head = iterNext

			} else if iter.Next.Next != nil {
				iterNext := iter.Next     
				iter.Next = iterNext.Next 
				iterNext.Next = iterNext.Next.Next
				iter.Next.Next = iterNext

				iter = iterNext       

			} else {
				break
			}

		} else {
			break
		}

	}

	return head
}