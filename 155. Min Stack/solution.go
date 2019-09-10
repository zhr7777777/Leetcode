type MinStack struct {
    Stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    var instance MinStack 
    return instance
}


func (this *MinStack) Push(x int)  {
    this.Stack = append(this.Stack, x)
}


func (this *MinStack) Pop()  {
    if len(this.Stack) != 0 {
        this.Stack = this.Stack[0:len(this.Stack)-1]
    }
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
    result := this.Stack[0]
    for i:=1; i<len(this.Stack); i++ {
        if result > this.Stack[i] {
            result = this.Stack[i]
        }
    }
    return result
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */