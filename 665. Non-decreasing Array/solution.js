/**
 * @param {number[]} nums
 * @return {boolean}
 */
var checkPossibility = function(nums) {
    if(nums.length === 1) {
        return true
    }
    let t1 = 0
    let t2 = 0
    let copy1 = [...nums]
    let copy2 = [...nums]
    // 分两种情况：1. 把前一项改小。2. 把当前项改大
    for(let i=1; i<copy1.length; i++) {
        if(copy1[i] < copy1[i-1]) {
            t1++
            copy1[i-1] = copy1[i]
            if(i-2 >= 0 && copy1[i-1] < copy2[i-2]) { // 把前一项改小时，注意前一项之前的序列保证依旧是升序
                t1++
            }
        }
    }
    for(let i=1; i<copy2.length; i++) {
        if(copy2[i] < copy2[i-1]) {
            t2++
            copy2[i] = copy2[i-1]
        }
    }
    return t1 < 2 || t2 < 2
};