/**
 * @param {number} n
 * @param {number} k
 * @return {string}
 */
var getPermutation = function(n, k) {
    let result = ''
    let arr = []
    for(let i=0; i<n; i++) {
        arr.push(i+1)
    }
    k--
    for(let i=0; i<n; i++) {
        let divider = factorial(n-1-i)
        let q = parseInt(k / divider)
        let r = k % divider
        result += String(arr[q])
        arr.splice(q, 1)
        k = r
    }
    return result
};

function factorial (num) { 
    if (num < 0) { 
        return -1; 
    } else if (num === 0 || num === 1) { 
        return 1; 
    } else { 
        return (num * factorial(num - 1)); 
    } 
}