/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func splitListToParts(head *ListNode, k int) []*ListNode {
    listNodes := make([]*ListNode, k)
	if head == nil {
		return listNodes
	}

	

	length := 0
	iter := head

	for iter != nil {
		length++
		iter = iter.Next
	}

	listNodes[0] = head
	elements := length / k
	//fmt.Printf("Elements : %d\n", elements)
	count := 1
	index := 1
	iter = head
	offset := 0

	if length > k {
		offset = length % k
	}
	//fmt.Printf("offset : %d\n", offset)

	for iter.Next != nil {
		nextH := iter.Next
		if count >= elements && index < len(listNodes) {

			if offset > 0 {
				iter = nextH
				nextH = nextH.Next
				offset--
			}

			fmt.Println(iter)
			listNodes[index] = nextH
			iter.Next = nil
			count = 1
			index++
		} else {
			count++
		}
		iter = nextH
	}

	return listNodes
}
