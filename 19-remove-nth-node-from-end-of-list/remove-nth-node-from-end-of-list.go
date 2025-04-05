/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if n == 1 && head.Next == nil {
		return nil
	}

	cursorEnd := head
	cursorPrev := head
	length := 1

	for cursorEnd.Next != nil {
		cursorPrev = cursorEnd
		cursorEnd = cursorEnd.Next
		length++
	}

	if n == 1 {
		cursorPrev.Next = nil
		return head
	} else if n == length {
		return head.Next
	}

	iter := head
	index := 1
	toFind := length - n

	for iter != nil {

		if index == toFind {
			iter.Next = iter.Next.Next
		}
		iter = iter.Next
		index++
	}

	return head
}