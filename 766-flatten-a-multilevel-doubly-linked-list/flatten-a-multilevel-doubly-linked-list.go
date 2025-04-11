/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {

	if root == nil {
		return nil
	}

	iter := root

	for iter != nil {
		if iter.Child != nil {
			var next *Node
			if iter.Next != nil {
				next = iter.Next

			}
			child := iter.Child
			childIter := child

			for childIter != nil {
				if childIter.Child != nil {
					_ = flatten(child)
					break
				}
				childIter = childIter.Next
			}

			iter.Next = child
			child.Prev = iter
			for child.Next != nil {
				child = child.Next
			}
			if next != nil {
				child.Next = next
				next.Prev = child
			} else {
				child.Next = nil
			}
			iter.Child = nil

		}

		iter = iter.Next
	}

	return root
}
