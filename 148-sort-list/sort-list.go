/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	middleNode := middle(head)

	left := sortList(head)
	right := sortList(middleNode)

	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {

	dumm := new(ListNode)
	dummEnd := dumm

	for left != nil && right != nil {

		if left.Val <= right.Val {
			dummEnd.Next = left
			left = left.Next
		} else {
			dummEnd.Next = right
			right = right.Next
		}

		dummEnd = dummEnd.Next
	}

	if left == nil {
		dummEnd.Next = right
	} else if right == nil {
		dummEnd.Next = left
	}

	return dumm.Next

}
func middle(head *ListNode) *ListNode {
	slow := head
	fast := head
	leadingSlow := slow
	for fast != nil && fast.Next != nil {
		leadingSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	leadingSlow.Next = nil
	return slow
}