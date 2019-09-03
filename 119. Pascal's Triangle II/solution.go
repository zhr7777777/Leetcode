var pascalTriangle [][]int

func getRow(rowIndex int) []int {
    if len(pascalTriangle) == 0 {
        var first = []int{1}
        pascalTriangle = append(pascalTriangle, first)   
    }
    if len(pascalTriangle) == 1 {
        var second = []int{1, 1}
        pascalTriangle = append(pascalTriangle, second)
    }
    if rowIndex < len(pascalTriangle) {
        return pascalTriangle[rowIndex]
    }
    curLength := len(pascalTriangle)
    for i:=0; i<rowIndex+1-curLength; i++ {
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
    return pascalTriangle[len(pascalTriangle)-1]
}