var pascalTriangle [][]int


func generate(numRows int) [][]int {
    if len(pascalTriangle) == 0 {
        var first = []int{1}
        pascalTriangle = append(pascalTriangle, first)   
    }
    if len(pascalTriangle) == 1 {
        var second = []int{1, 1}
        pascalTriangle = append(pascalTriangle, second)
    }
    if numRows <= len(pascalTriangle) {
        return pascalTriangle[:numRows]
    }
    curLength := len(pascalTriangle)
    for i:=0; i<numRows-curLength; i++ {
        nextRow := []int{1}
        for j:=0; j<len(pascalTriangle[len(pascalTriangle)-1]); j++ {
            if j + 1 == len(pascalTriangle[len(pascalTriangle)-1]) {
                break
            }
            nextRow = append(nextRow, pascalTriangle[len(pascalTriangle)-1][j] + pascalTriangle[len(pascalTriangle)-1][j+1])
        }
        nextRow = append(nextRow, 1)
        pascalTriangle = append(pascalTriangle, nextRow)
    }
    return pascalTriangle
}