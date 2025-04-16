/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var pos *ListNode

	for head != nil {
		nextH := head.Next
		head.Next = pos
		pos = head
		head = nextH

	}

	return pos
}

func doubleIt(head *ListNode) *ListNode {
	if head.Next == nil {
		double := head.Val * 2
		if double >= 10 {
			head.Val = (head.Val * 2) % 10
			return &ListNode{1, head}
		} else {
			head.Val = (head.Val * 2)
			return head
		}

	}

	head = reverseList(head)
	iter := head
	rem := 0

	for iter != nil {
		doub := iter.Val * 2
		iter.Val = (doub % 10) + rem
		if doub >= 10 {
			rem = 1
		} else {
			rem = 0
		}
		iter = iter.Next
	}

	head = reverseList(head)
	if rem == 1 {
		return &ListNode{1, head}
	}
	return head

}