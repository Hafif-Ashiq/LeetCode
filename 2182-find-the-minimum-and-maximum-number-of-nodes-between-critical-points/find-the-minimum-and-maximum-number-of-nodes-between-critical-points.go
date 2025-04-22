/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	answer := []int{-1, -1}

	if head.Next.Next == nil {
		return answer
	}

	iter := head.Next
	leading := head
	i := 2

	first := 0
	prev := 0
	last := 0

	for iter.Next != nil {
		isCritical := false
		if (iter.Val > leading.Val && iter.Val > iter.Next.Val) || (iter.Val < leading.Val && iter.Val < iter.Next.Val) {
			
			isCritical = true
		}
		if isCritical {
			if first == 0 {
				first = i
			}
			prev = last
			last = i

			if prev != 0 {
				diff := last - prev
				if answer[0] == -1 || diff < answer[0] {
					answer[0] = diff
				}
			}
		}

		leading = iter
		iter = iter.Next
		i++
	}
	if last != 0 && first != 0 && answer[0] != -1 {
		answer[1] = last - first
	}
	return answer

}