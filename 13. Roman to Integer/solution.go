func romanToInt(s string) int {
    romanToIntMap := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    if(len(s) == 1) {
        return romanToIntMap[s]
    }
    var result = 0
    for i:=0; i<len(s); i++ {
        var cur int = romanToIntMap[string(s[i])]
        if(i + 1 == len(s)) {
            result += cur
            break
        }
        var next int = romanToIntMap[string(s[i+1])]
        if(cur < next) {
            result -= cur
        } else {
            result += cur
        }
    }
    return result
}