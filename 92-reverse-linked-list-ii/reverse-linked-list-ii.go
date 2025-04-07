/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	last := head

	for last.Next != nil {
		last = last.Next
	}

	anchorL := last

	for {
		iter := head

		for iter.Next != anchorL {
			iter = iter.Next
		}
		anchorL.Next = iter
		anchorL = iter

		if anchorL == head {
			break
		}
	}
	head.Next = nil
	return last
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil || head.Next == nil || left == right {
		return head
	}

	index := 1
	iter := head

	var beforeLeft *ListNode
	var afterRight *ListNode

	var reverseStart *ListNode
	var reverseEnd *ListNode

	for iter != nil {

		if index == left-1 {
			beforeLeft = iter

		} else if index == right+1 {
			afterRight = iter
		} else if index == left {
			reverseStart = iter
		} else if index == right {
			reverseEnd = iter
		}

		index++
		iter = iter.Next
	}

	reverseEnd.Next = nil
	reversed := reverseList(reverseStart)

	reverseStart.Next = afterRight

	if beforeLeft == nil {
		return reversed
	}
	beforeLeft.Next = reversed

	return head
}