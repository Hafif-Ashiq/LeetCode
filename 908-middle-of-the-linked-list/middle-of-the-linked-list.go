/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
	slow := head
	fast := head
	leadingSlow := slow
	for fast != nil && fast.Next != nil {
		leadingSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	leadingSlow.Next = nil
	return slow
}