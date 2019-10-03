type MyStack struct {
    Array []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    var stack MyStack
    stack.Array = []int{}
    return stack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Array = append(this.Array, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    top := 0
    temp := []int{}
    length := len(this.Array)
    for i:=0; i<length; i++ {
        if len(this.Array) == 1 {
            top = this.Array[0]
            this.Array = temp
            break
        }
        temp = append(temp, this.Array[0])
        this.Array = this.Array[1:len(this.Array)]
    }
    return top
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.Array[len(this.Array)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.Array) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */