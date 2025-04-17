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

func middle(head *ListNode) *ListNode {

	slow := head
	fast := head
	leading := slow
	for fast != nil {

		leading = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	leading.Next = nil
	return slow
}

func pairSum(head *ListNode) int {

	if head.Next.Next == nil {
		return head.Val + head.Next.Val
	}

	secondH := middle(head)
	fir := head
	sec := reverseList(secondH)

	max := 0
	for fir != nil {
		if fir.Val+sec.Val > max {
			max = fir.Val + sec.Val
		}
		fir = fir.Next
		sec = sec.Next
	}
	return max

}