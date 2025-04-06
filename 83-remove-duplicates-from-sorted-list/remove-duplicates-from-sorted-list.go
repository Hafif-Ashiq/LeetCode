/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	iter := head

	for iter != nil {
		curr := iter
		for curr.Next != nil && curr.Next.Val == iter.Val {
			curr = curr.Next
		}
		iter.Next = curr.Next
		iter = iter.Next

	}

	return head
}