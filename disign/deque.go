package disign

// The size of a block of data
const blockSize = 4096

// Double ended queue data structure.
type Deque struct {
	leftIdx  int
	leftOff  int
	rightIdx int
	rightOff int

	blocks [][]interface{}
	left   []interface{}
	right  []interface{}
}

// Creates a new, empty deque.
func New() *Deque {
	result := new(Deque)
	result.blocks = [][]interface{}{make([]interface{}, blockSize)}
	result.right = result.blocks[0]
	result.left = result.blocks[0]
	return result
}

// Pushes a new element into the queue from the right, expanding it if necessary.
func (d *Deque) PushRight(data interface{}) {
	d.right[d.rightOff] = data
	d.rightOff++
	if d.rightOff == blockSize {
		d.rightOff = 0
		d.rightIdx = (d.rightIdx + 1) % len(d.blocks)

		// If we wrapped over to the left, insert a new block and update indices
		if d.rightIdx == d.leftIdx {
			buffer := make([][]interface{}, len(d.blocks)+1)
			copy(buffer[:d.rightIdx], d.blocks[:d.rightIdx])
			buffer[d.rightIdx] = make([]interface{}, blockSize)
			copy(buffer[d.rightIdx+1:], d.blocks[d.rightIdx:])
			d.blocks = buffer
			d.leftIdx++
			d.left = d.blocks[d.leftIdx]
		}
		d.right = d.blocks[d.rightIdx]
	}
}

// Pops out an element from the queue from the right. Note, no bounds checking are done.
func (d *Deque) PopRight() (res interface{}) {
	d.rightOff--
	if d.rightOff < 0 {
		d.rightOff = blockSize - 1
		d.rightIdx = (d.rightIdx - 1 + len(d.blocks)) % len(d.blocks)
		d.right = d.blocks[d.rightIdx]
	}
	res, d.right[d.rightOff] = d.right[d.rightOff], nil
	return
}

// Returns the rightmost element from the deque. No bounds are checked.
func (d *Deque) Right() interface{} {
	if d.rightOff > 0 {
		return d.right[d.rightOff-1]
	} else {
		return d.blocks[(d.rightIdx-1+len(d.blocks))%len(d.blocks)][blockSize-1]
	}
}

// Pushes a new element into the queue from the left, expanding it if necessary.
func (d *Deque) PushLeft(data interface{}) {
	d.leftOff--
	if d.leftOff < 0 {
		d.leftOff = blockSize - 1
		d.leftIdx = (d.leftIdx - 1 + len(d.blocks)) % len(d.blocks)

		// If we wrapped over to the right, insert a new block and update indices
		if d.leftIdx == d.rightIdx {
			d.leftIdx++
			buffer := make([][]interface{}, len(d.blocks)+1)
			copy(buffer[:d.leftIdx], d.blocks[:d.leftIdx])
			buffer[d.leftIdx] = make([]interface{}, blockSize)
			copy(buffer[d.leftIdx+1:], d.blocks[d.leftIdx:])
			d.blocks = buffer
		}
		d.left = d.blocks[d.leftIdx]
	}
	d.left[d.leftOff] = data
}

// Pops out an element from the queue from the left. Note, no bounds checking are done.
func (d *Deque) PopLeft() (res interface{}) {
	res, d.left[d.leftOff] = d.left[d.leftOff], nil
	d.leftOff++
	if d.leftOff == blockSize {
		d.leftOff = 0
		d.leftIdx = (d.leftIdx + 1) % len(d.blocks)
		d.left = d.blocks[d.leftIdx]
	}
	return
}

// Returns the leftmost element from the deque. No bounds are checked.
func (d *Deque) Left() interface{} {
	return d.left[d.leftOff]
}

// Checks whether the queue is empty.
func (d *Deque) Empty() bool {
	return d.leftIdx == d.rightIdx && d.leftOff == d.rightOff
}

// Returns the number of elements in the queue.
func (d *Deque) Size() int {
	if d.rightIdx > d.leftIdx {
		return (d.rightIdx-d.leftIdx)*blockSize - d.leftOff + d.rightOff
	} else if d.rightIdx < d.leftIdx {
		return (len(d.blocks)-d.leftIdx+d.rightIdx)*blockSize - d.leftOff + d.rightOff
	} else {
		return d.rightOff - d.leftOff
	}
}

// Clears out the contents of the queue.
func (d *Deque) Reset() {
	*d = *New()
}

type MyCircularDeque1 struct {
	head *Node
	tail *Node
	len  int
	cap  int
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor1(k int) MyCircularDeque1 {
	head := Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	tail := Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	head.Next = &tail
	tail.Pre = &head

	return MyCircularDeque1{
		head: &head,
		tail: &tail,
		len:  0,
		cap:  k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque1) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.head.Next
	tempNode := Node{
		Val:  value,
		Next: temp,
		Pre:  this.head,
	}
	this.head.Next = &tempNode
	temp.Pre = &tempNode
	this.len++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque1) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.tail.Pre
	tempNode := Node{
		Val:  value,
		Next: this.tail,
		Pre:  temp,
	}
	this.tail.Pre = &tempNode
	temp.Next = &tempNode
	this.len++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque1) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.head.Next
	this.head.Next = deleteTemp.Next
	deleteTemp.Next.Pre = this.head
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque1) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.tail.Pre
	this.tail.Pre = deleteTemp.Pre
	deleteTemp.Pre.Next = this.tail
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.len--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque1) GetFront() int {
	return this.head.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque1) GetRear() int {
	return this.tail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque1) IsEmpty() bool {
	return this.head.Next == this.tail
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque1) IsFull() bool {
	return this.len == this.cap
}
