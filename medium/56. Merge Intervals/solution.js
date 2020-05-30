/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function(intervals) {
    const compareIntervals = (a, b) => {
        if(a[0] > b[0]) {
            return 1
        }
        if(a[0] < b[0]) {
            return -1
        }
        return 0
    }
    intervals = intervals.sort(compareIntervals)
    let length = intervals.length
    for(let i=0; i<length; i++) {
        if(i+1 < length) {
            if(intervals[i+1][0] <= intervals[i][1]) {
                intervals[i][1] = Math.max(intervals[i][1], intervals[i+1][1])
                intervals.splice(i+1, 1) // ！！！注意使用splice时数组长度会变，循环之前应该记录数组长度
                // intervals[i+1] = intervals[intervals.length-1] // 不能这样，破坏了排序
                // intervals.length--
                length--    // splice后数组长度减一，防止数组越界
                i--         // 合并了后一区间，下一次应再从当前区间开始判断
            }
        }
    }
    return intervals
};