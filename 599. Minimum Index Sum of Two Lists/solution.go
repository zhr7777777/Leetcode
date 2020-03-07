func findRestaurant(list1 []string, list2 []string) []string {
    mapping := map[string]int{}
    result := []string{}
    curMin := len(list1) + len(list2)
    for i:=0; i<len(list1); i++ {
        mapping[list1[i]] = i
    }
    for i:=0; i<len(list2); i++ {
        index, has := mapping[list2[i]]
        if has {
            if index + i < curMin {
                result = []string{}
                result = append(result, list2[i])
                curMin = index + i
            } else if index + i == curMin {
                result = append(result, list2[i])
            }
        }
    }
    return result
}