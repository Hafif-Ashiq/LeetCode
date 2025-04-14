/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

	if head == nil {
		return head
	}

	newList := new(Node)
	newEnd := newList

	iter := head

	randomMap := make(map[*Node]*Node)

	for iter != nil {
		newEnd.Next = &Node{iter.Val, nil, iter.Random}
		// Store addresses
		randomMap[iter] = newEnd.Next
		iter = iter.Next
		newEnd = newEnd.Next
	}

	newList = newList.Next

	iter = newList

	for iter != nil {
		if iter.Random != nil {
			iter.Random = randomMap[iter.Random]
		}
		iter = iter.Next
	}

	return newList

}