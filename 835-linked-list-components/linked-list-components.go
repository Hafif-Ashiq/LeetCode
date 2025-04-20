/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func numComponents(head *ListNode, nums []int) int {
	if head.Next == nil && slices.Contains(nums, head.Val) {
		return 1
	}

	count := 0

	set := make(map[int]bool, len(nums))
	for _, v := range nums {
		set[v] = true
	}

	iter := head

	nonCons := false
	for iter.Next != nil {
		if iter == head && (set[iter.Val] && !set[iter.Next.Val]) {
			// fmt.Println("Here")
			// fmt.Println(iter.Val)
			count++
		} else if set[iter.Val] && set[iter.Next.Val] {
			// fmt.Println("Here 2")
			// fmt.Println(iter.Val)
			// h2 := iter.Next
			for iter.Next.Next != nil && set[iter.Next.Val] {
				iter = iter.Next
			}

			count++
			nonCons = false
		} else if nonCons && (set[iter.Val]) {
			// fmt.Println("Here 3")
			// fmt.Println(iter.Val)
			count++
			nonCons = false
		} else {
			nonCons = true
		}
		iter = iter.Next
	}
	if set[iter.Val] && nonCons {
		count++
	}
	return count

}