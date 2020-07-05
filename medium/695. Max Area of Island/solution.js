/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxAreaOfIsland = function(grid) {
    const count = (i, j) => {
        if(i<0 || i>=grid.length || j<0 || j>=grid[i].length || grid[i][j] === 0) return 0
        grid[i][j] = 0
        return 1 + count(i-1, j) + count(i+1, j) + count(i, j-1) + count(i, j+1)
    }
    let max = 0
    for(let i=0; i<grid.length; i++) {
        for(let j=0; j<grid[i].length; j++) {
            if(grid[i][j] === 1) {
               max = Math.max(max, count(i, j)) 
            }
        }
    }
    return max
};