var slice = []int{1, 2, 3}

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    if n == 3 {
        return 3
    }
    if n-3 == len(slice) - 1 {
        next := slice[n-3] + slice[n-4]
        last := slice[n-3]
        slice = append(slice, next)
        slice = append(slice, last + next)
        return last + next
    }
    if n-2 == len(slice) - 1 {
        next := slice[n-2] + slice[n-3]
        slice = append(slice, next)
        return next
    }
    if n-2 < len(slice) - 1 {
        return slice[n-1]
    }
    return climbStairs(n-1) + climbStairs(n-2)
}