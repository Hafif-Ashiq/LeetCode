/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	prevMiddle, middle := middleNode(head)

	prevMiddle.Next = middle.Next

	return head
}

func middleNode(head *ListNode) (*ListNode, *ListNode) {
	slow := head
	fast := head
	leadingSlow := slow
	for fast != nil && fast.Next != nil {
		leadingSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return leadingSlow, slow
}
