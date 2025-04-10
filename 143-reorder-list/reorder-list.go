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

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	length := 1
	last := head

	for last.Next != nil {
		last = last.Next
		length++
	}

	half := length / 2

	before := new(ListNode)
	after := new(ListNode)

	beforeEnd := before
	afterEnd := after
	len := 1

	for head != nil {
		if len <= half {
			beforeEnd.Next = head
			beforeEnd = head
		} else {
			afterEnd.Next = head
			afterEnd = head
		}
		head = head.Next
		len++
	}
	beforeEnd.Next = nil
	afterEnd.Next = nil

	final := new(ListNode)
	finalEnd := final
	after = reverseList(after.Next)
	before = before.Next

	for after != nil || before != nil {
		if before != nil {
			finalEnd.Next = before
			finalEnd = before
			before = before.Next
		}
		if after != nil {
			finalEnd.Next = after
			finalEnd = after
			after = after.Next
		}
	}

}