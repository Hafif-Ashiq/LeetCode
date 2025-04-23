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
		nextH := head.Next
		head.Next = pos
		pos = head
		head = nextH
	}

	return pos
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {

	curr := 1
	i := 0

	iter := head
	leading := head

	for iter != nil {

		count := 0
		nextIter := iter
		var leadingNext *ListNode
		
		for count < curr && nextIter != nil {
			leadingNext = nextIter
			nextIter = nextIter.Next
			count++
		}
		
        if  count%2 == 0 {
		
        	leadingNext.Next = nil
			rev := reverseList(iter)
		
        	leading.Next = rev
			leading = iter
			iter.Next = nextIter
			iter = iter.Next
			i = 0
		} else {
			leading = leadingNext
			iter = nextIter
			i++
		}
        curr++


	}

	return head

}