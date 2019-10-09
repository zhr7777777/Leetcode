func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapping := map[string]int{}
    for i:=0; i<len(s); i++ {
        c := string(s[i])
        count, has := mapping[c]
        if has {
            mapping[c] = count + 1
        } else {
            mapping[c] = 1
        }
    }
    for i:=0; i<len(t); i++ {
        c := string(t[i])
        count, has := mapping[c]
        if has {
            mapping[c] = count - 1
        } else {
            return false
        }
    }
    for c := range mapping {
        if mapping[c] != 0 {
            return false
        }
    }
    return true
}

// func isAnagram(s string, t string) bool {
//     if len(s) != len(t) {
//         return false
//     }
//     if sortString(s) == sortString(t) {
//         return true
//     }
//     return false
// }

// func sortString(s string) string {
//     for i:=0; i<len(s)-1; i++ {
//         for j:=0; j<len(s)-i-1; j++ {
//             a := []rune(s)[j]
//             b := []rune(s)[j+1]
//             if a > b {
//                 temp := b
//                 s = replaceAtIndex(s, a, j+1)
//                 s = replaceAtIndex(s, temp, j)
//             }
//         }
//     }
//     return s
// }

// func replaceAtIndex(in string, r rune, i int) string {
//     out := []rune(in)
//     out[i] = r
//     return string(out)
// }