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

func removeNodes(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}

	dumm := new(ListNode)
	dummEnd := dumm

	head = reverseList(head)
	max := 0
	for head != nil {
		if head.Val >= max {
			dummEnd.Next = head
			dummEnd = dummEnd.Next
			max = head.Val

		}
		head = head.Next
	}

	dummEnd.Next = nil
	dumm = dumm.Next

	return reverseList(dumm)

}