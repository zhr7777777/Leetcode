type MyQueue struct {
    Array []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    var stack MyQueue
    stack.Array = []int{}
    return stack
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.Array = append(this.Array, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    first := this.Array[0]
    this.Array = this.Array[1:]
    return first
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.Array[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.Array) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */