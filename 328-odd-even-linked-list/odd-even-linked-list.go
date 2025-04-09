/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	dummOdd := &ListNode{-1, nil}
	dummEven := &ListNode{-1, nil}

	oddTail := dummOdd
	evenTail := dummEven

	iter := head

	index := 1

	for iter != nil {
		if index%2 == 0 {
			evenTail.Next = iter
			evenTail = iter
			iter = iter.Next
			evenTail.Next = nil
		} else {
			oddTail.Next = iter
			oddTail = iter
			iter = iter.Next
			oddTail.Next = nil
		}
		index++
	}

	oddTail.Next = dummEven.Next

	return dummOdd.Next
}
