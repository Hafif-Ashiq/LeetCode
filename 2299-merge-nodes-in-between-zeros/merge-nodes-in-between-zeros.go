/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {

	final := new(ListNode)
	finalEnd := final

	iter := head.Next

	sum := 0
	for iter != nil {

		if iter.Val == 0 {
			finalEnd.Next = &ListNode{sum, nil}
			finalEnd = finalEnd.Next
			sum = 0

		} else {
			sum += iter.Val
		}
		iter = iter.Next
	}

	return final.Next

}