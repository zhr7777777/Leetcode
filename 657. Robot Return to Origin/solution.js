/**
 * @param {string} moves
 * @return {boolean}
 */
var judgeCircle = function(moves) {
    if(moves.length % 2 === 1) {
        return false
    }
    let rights = 0
    let lefts = 0
    let ups = 0
    let downs = 0
    for(let i=0; i<moves.length; i++) {
        switch(moves.charAt(i)) {
            case 'R':
                rights++;break;
            case 'L':
                lefts++;break;
            case 'U':
                ups++;break;
            default:
                downs++;break;
        }
    }
    return rights === lefts && ups === downs
};