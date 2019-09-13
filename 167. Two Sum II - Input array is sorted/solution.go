func twoSum(numbers []int, target int) []int {
    var numsMap map[int]int
    var result [] int
    numsMap = make(map[int]int)
    for i := 0; i<len(numbers); i++ {
        search := 0
        if target < numbers[len(numbers) - 1] {
            search = target - numbers[i]
        } else {
            search = target - numbers[len(numbers)-1-i]
        }
        idx, has := numsMap[search]
        if has {
            if target < numbers[len(numbers) - 1] { // important condition
                result = append(result, idx+1)
                result = append(result, i+1)
            } else {
                result = append(result, len(numbers)-1-i+1)
                result = append(result, idx+1)
            }
            break
        }
        if target < numbers[len(numbers) - 1] {
            numsMap[numbers[i]] = i
        } else {
            numsMap[numbers[len(numbers)-1-i]] = len(numbers)-1-i
        }
    }
    return result
}