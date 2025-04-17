/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func gcd(a, b int) int {
	// If b is 0, the GCD is a
	if b == 0 {
		return a
	}
	// Recursively call gcd with b and the remainder of a divided by b
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	iter := head

	for iter.Next != nil {
		gcdV := gcd(iter.Val, iter.Next.Val)
		next := iter.Next
		iter.Next = &ListNode{gcdV, next}
		iter = next
	}
	return head
}
