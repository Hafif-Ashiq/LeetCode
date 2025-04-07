/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var before *ListNode
	var beforeLast *ListNode

	var after *ListNode
	var afterLast *ListNode

	iter := head

	for iter != nil {

		if iter.Val < x {
			fmt.Println(iter.Val)
			if before == nil {
				before = iter
				beforeLast = iter
			} else {
				beforeLast.Next = iter
				beforeLast = iter
			}
		} else {
			if after == nil {
				after = iter
				afterLast = iter
			} else {
				afterLast.Next = iter
				afterLast = iter
			}
		}
		iter = iter.Next
	}

    if before == nil {
		return after
	}
	beforeLast.Next = after
	if afterLast != nil {
		afterLast.Next = nil
	}

	return before

}