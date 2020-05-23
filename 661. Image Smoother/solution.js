/**
 * @param {number[][]} M
 * @return {number[][]}
 */
var imageSmoother = function(M) {
    let result = JSON.parse(JSON.stringify(M))
    for(let i=0; i<M.length; i++) {
        for(let j=0; j<M[i].length; j++) {
            let sum = M[i][j]
            let count = 1
            if(j+1 < M[i].length) {
                sum += M[i][j+1]
                count++
            }
            if(i+1 < M.length) {
                sum += M[i+1][j]  
                count++
            }
            if(j+1 < M[i].length && i+1 < M.length) {
                sum += M[i+1][j+1] 
                count++
            }
            if(j+1 < M[i].length && i-1 >= 0) {
                sum += M[i-1][j+1] 
                count++
            }
            if(j-1 >=0 && i+1 < M.length) {
                sum += M[i+1][j-1] 
                count++
            }
            if(j-1 >= 0) {
                sum += M[i][j-1]
                count++
            }
            if(i-1 >= 0) {
                sum += M[i-1][j]
                count++
            }
            if(j-1 >= 0 && i-1 >= 0) {
                sum += M[i-1][j-1]
                count++
            }
            result[i][j] = Math.floor(sum / count)
        }
    }
    return result
};