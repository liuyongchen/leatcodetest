package disign

import (
	"container/list"
)

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

type CQueue1 struct {
	in  stack
	out stack
}

type stack []int

func (s *stack) Push(value int) {
	*s = append(*s, value)
}

func (s *stack) Pop() int {
	n := len(*s)
	res := (*s)[n-1]

	*s = (*s)[:n-1]

	return res
}

func Constructor2() CQueue {
	return CQueue{}
}

func (this *CQueue1) AppendTail2(value int) {
	this.in.Push(value)
}

func (this *CQueue1) DeleteHead2() int {
	if len(this.out) != 0 {
		return this.out.Pop()
	} else if len(this.in) != 0 {
		for len(this.in) != 0 {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop()
	}

	return -1

}
