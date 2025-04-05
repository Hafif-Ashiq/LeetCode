/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	start1 := l1
	start2 := l2
	remainder := 0

	newList := &ListNode{0, nil}
	start3 := newList

	for {

		sum := 0
		if start1 == nil && start2 == nil {

			break
		} else if start1 == nil {
			sum = start2.Val
		} else if start2 == nil {
			sum = start1.Val
		} else {
			sum = start1.Val + start2.Val
		}

		newSum := remainder + (sum % 10)

		if sum > 9 {
			remainder = 1
		} else {
			remainder = 0
		}

		if newSum == 10 {
			newSum = 0
			remainder = 1
		}

		start3.Val = newSum

		if start1 != nil {
			start1 = start1.Next
		}
		if start2 != nil {
			start2 = start2.Next
		}

		if start1 != nil || start2 != nil {

			start3.Next = &ListNode{}
			start3 = start3.Next

		} else {
			if remainder != 0 {
				start3.Next = &ListNode{}
				start3 = start3.Next
				start3.Val = remainder
			}
		}

	}
	start3.Next = nil

	return newList

}
