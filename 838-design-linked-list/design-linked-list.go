package main

import "fmt"

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func Constructor() MyLinkedList {

 	return MyLinkedList{nil, nil, 0}
}

func (this *MyLinkedList) Get(index int) int {
	i := 0
	iter := this.Head
	for i != index && iter != nil {
		iter = iter.Next
		i++
	}
	if iter == nil {
		return -1
	}
	return iter.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.Size++
	if this.Head == nil {
		this.Head = &Node{val, nil}
		this.Tail = this.Head
		return
	}

	new := &Node{val, this.Head}
	this.Head = new

}

func (this *MyLinkedList) AddAtTail(val int) {
	this.Size++
	if this.Head == nil {
		this.Head = &Node{val, nil}
		this.Tail = this.Head
		return
	}
	new := &Node{val, nil}
	this.Tail.Next = new
	this.Tail = new

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index < 0 || index > this.Size {
		// Index out of bounds, do nothing
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.Size {
		this.AddAtTail(val)
		return
	}

	i := 0
	iter := this.Head
	leading := this.Head

	for i < index && iter != nil {
		leading = iter
		iter = iter.Next
		i++
	}

	if iter == nil || i != index {
		return
	}

	this.Size++
	leading.Next = &Node{val, iter}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 && this.Head != nil {
		this.Head = this.Head.Next
		this.Size--
		return
	}
	i := 0
	iter := this.Head
	leading := this.Head

	for i != index && iter != nil {
		leading = iter
		iter = iter.Next
		i++
	}

	if iter == nil {
		return
	}

	if iter == this.Tail {
		this.Tail = leading
	}
	leading.Next = iter.Next
	this.Size--
}