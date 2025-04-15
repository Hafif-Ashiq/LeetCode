/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	dumm := new(ListNode)
	dummEnd := dumm
	set := make(map[int]struct{})
	emt := struct{}{}
	for _, v := range nums {
		set[v] = emt
	}

	iter := head

	for iter != nil {
		_, ok := set[iter.Val]
		if !ok {
			dummEnd.Next = iter
			dummEnd = dummEnd.Next
		}
		iter = iter.Next
	}
	dummEnd.Next = nil

	return dumm.Next
}