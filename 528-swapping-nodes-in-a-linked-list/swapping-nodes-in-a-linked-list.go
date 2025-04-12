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

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	iter := head
	var firstNode *ListNode
	var secondNode *ListNode

	ind := 1
	for iter != nil {
		if ind == k {
			firstNode = iter
			break
		}
		iter = iter.Next
		ind++
	}

	head = reverseList(head)
	iter = head

	ind = 1
	for iter != nil {
		if ind == k {
			secondNode = iter
			break
		}
		iter = iter.Next
		ind++

	}

	head = reverseList(head)

	if firstNode == nil || secondNode == nil {
		return head
	}

	temp := firstNode.Val
	firstNode.Val = secondNode.Val
	secondNode.Val = temp

	return head

}
