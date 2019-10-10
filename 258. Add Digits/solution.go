// func addDigits(num int) int {
//     if num < 10 {
//         return num
//     }
//     result := 0
//     for {
//         if num == 0 {
//             break
//         }
//         result = result + (num % 10)
//         num = int(num / 10)
//     }
//     return addDigits(result)
// }

func addDigits(num int) int {
    if num == 0 {
        return 0
    }
    if num % 9 == 0 {
        return 9
    }
    return num % 9
}