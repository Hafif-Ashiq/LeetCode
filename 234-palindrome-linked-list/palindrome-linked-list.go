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

func isPalindrome(head *ListNode) bool {

	if head.Next == nil {
		return true
	}

	iter := head
	arr := make([]int, 0)

	for iter != nil {
		arr = append(arr, iter.Val)
		iter = iter.Next
	}

	head = reverseList(head)

	iter = head
	i := 0
	for iter != nil {
		if arr[i] != iter.Val {
			return false
		}
		iter = iter.Next
        i++
	}

	return true

}