/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "math/rand/v2"
type Solution struct {
	head   *ListNode
	length int
}

func Constructor(head *ListNode) Solution {

	leng := 0
	iter := head
	for iter != nil {
		leng++
		iter = iter.Next
	}

	return Solution{head, leng}
}

func (this *Solution) GetRandom() int {

	rand := rand.IntN(this.length)

	i := 0
	iter := this.head
	for i < rand {
		iter = iter.Next
		i++
	}

	return iter.Val

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */