/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {

	dec := 0
	length := 0
	iter := head
	for iter != nil {
		length++
		iter = iter.Next
	}

	length--
	for head != nil {
		dec += head.Val * int(math.Pow(2, float64(length)))
		length--
		head = head.Next
	}

	return dec
}
