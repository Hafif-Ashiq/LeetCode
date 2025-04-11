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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	} else if l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	}

	l1 = reverseList(l1)
	l2 = reverseList(l2)

	dumm := new(ListNode)
	dumIter := dumm
	rem := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum+rem > 9 {
			newVal := (sum + rem) % 10
			dumIter.Next = &ListNode{newVal, nil}
			dumIter = dumIter.Next
			rem = 1
		} else {
			dumIter.Next = &ListNode{sum + rem, nil}
			dumIter = dumIter.Next
			rem = 0
		}

	}

    if rem == 1{
        dumIter.Next =  &ListNode{1, nil}
    }

	return reverseList(dumm.Next)

}