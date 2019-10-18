type NumArray struct {
    Data []int
}


func Constructor(nums []int) NumArray {
    var instance NumArray
    instance.Data = nums
    return instance
}


func (this *NumArray) SumRange(i int, j int) int {
    sum := 0
    for k:=i; k<=j; k++ {
        sum += this.Data[k]
    }
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */