func isPalindrome(x int) bool {
    if(x < 0) {
        return false
    }
    if(x == 0) {
        return true
    }
    if(x % 10 == 0) {
        return false
    }
    var result int = 0
    var origin int = x
    for {
        if(x == 0) {
            break
        }
        var cur int = x % 10
        x = int(x / 10)
        result = result * 10 + cur
    }
    if(origin == result) {
        return true
    }
    return false
}