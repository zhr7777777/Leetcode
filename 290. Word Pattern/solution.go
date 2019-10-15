func wordPattern(pattern string, str string) bool {
    arr := strings.Split(str, " ")
    if len(arr) != len(pattern) {
        return false
    }
    mapping := map[string]string{}
    reverseMapping := map[string]string{}
    for i:=0; i<len(pattern); i++ {
        subStr, has := mapping[string(pattern[i])]
        if has {
            if subStr != arr[i] {
                return false
            } 
        } else {
            mapping[string(pattern[i])] = arr[i]
            
        }
    }
    for i:=0; i<len(arr); i++ { // important, similar to 205. Isomorphic Strings
        arrItem, has := reverseMapping[arr[i]]
        if has {
            if arrItem != string(pattern[i]) {
                return false
            } 
        } else {
            reverseMapping[arr[i]] = string(pattern[i])
        }
    }
    return true
}