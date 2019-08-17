func plusOne(digits []int) []int {
    for i:=0; i<len(digits); i++ {
        digits[len(digits)-1-i]++
        if digits[len(digits)-1-i] == 10 {
            digits[len(digits)-1-i] = 0
            if i == len(digits)-1 {
                var first = []int{1}
                digits = append(first, digits...)
                break
            }
        } else {
            break
        }
    }
    return digits
}