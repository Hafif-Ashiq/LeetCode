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
		nextH := head.Next // 2
		head.Next = pos
		pos = head // 1
		head = nextH

	}

	return pos
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	length := 1
	last := head

	for last.Next != nil {
		last = last.Next
		length++
	}

	groups := length / k

	final := new(ListNode)
	finalEnd := final

	iter := head
	for i := 0; i < groups; i++ {
		count := 0
		start := iter
		leading := iter
		for count < k {
			leading = iter
			iter = iter.Next
			count++
		}
		leading.Next = nil
		finalEnd.Next = reverseList(start)
		finalEnd = start
	}
	if iter != nil {
		finalEnd.Next = iter
	}
	return final.Next

}