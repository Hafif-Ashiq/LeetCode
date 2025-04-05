/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var finalList *ListNode

	head1 := list1
	head2 := list2

	var finalHead *ListNode

	for {
		if head1 == nil {
			finalHead.Next = head2
			break
		} else if head2 == nil {
			finalHead.Next = head1
			break
		} else {
			if head1.Val <= head2.Val {
				if finalHead == nil {
					finalList = head1
					finalHead = head1
					head1 = head1.Next
					finalHead.Next = nil
				} else {
					finalHead.Next = head1
					head1 = head1.Next
					finalHead = finalHead.Next
					finalHead.Next = nil
				}
			} else {
				if finalHead == nil {
					finalList = head2
					finalHead = head2
					head2 = head2.Next
					finalHead.Next = nil
				} else {
					finalHead.Next = head2
					head2 = head2.Next
					finalHead = finalHead.Next
					finalHead.Next = nil
				}
			}
		}
		if head1 == nil && head2 == nil {
			break
		}

	}

	return finalList
}