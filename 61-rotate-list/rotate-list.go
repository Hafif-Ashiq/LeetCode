/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	if k == 0 {
		return head
	}
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	last := head
	length := 1

	for last.Next != nil {
		last = last.Next
		length++
	}

	iter := head

	i := 0
	rotations := k % length

	if rotations == 0 {
		return head
	}
	var leading *ListNode
	for i < length-rotations {
		leading = iter

		iter = iter.Next
		i++
	}
	leading.Next = nil
	last.Next = head

	return iter

}