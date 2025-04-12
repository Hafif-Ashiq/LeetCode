/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	iterators := make(map[*ListNode]int)

	iter := headA

	for iter != nil {
		iterators[iter] = iter.Val
		iter = iter.Next
	}

	iter = headB

	for iter != nil {
		_, exist := iterators[iter]
		if exist {
			return iter
		}
		iter = iter.Next
	}

	return nil
}