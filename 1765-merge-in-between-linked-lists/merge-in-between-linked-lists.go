/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	iter := list1
	leading := list1
	i := 0
	var first *ListNode
	var second *ListNode

	for iter != nil {
		if i == a {
			first = leading
		} 
        if i == b {
			second = iter.Next
            break
		}
		leading = iter
		iter = iter.Next
		i++
	}

	if first != nil {
		first.Next = list2
	}
	iter = list2

	for iter.Next != nil {
		iter = iter.Next
	}

	iter.Next = second

	return list1

}