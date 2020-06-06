/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var multiply = function(num1, num2) {
    if(num1 === '0' || num2 === '0') return '0'
    let mul = '0'
    // if(num1.length > num2.length) {
    //     let temp = num1
    //     num1 = num2
    //     num2 = temp
    // }
    for(let i=0; i<num1.length; i++) {
        // mul = mul + Number(num1[num1.length-1-i]) * Number(num2) * Math.pow(10, i)
        // let cur = String(Number(num1[num1.length-1-i]) * Number(num2)) + '0'.repeat(i)
        let cur = mulStrings(num2, num1[num1.length-1-i]) + '0'.repeat(i)
        mul = addStrings(mul, cur)
    }
    return mul
};

var mulStrings = function (num1, num2) {
    let flag = 0
    let res = ''
    for(let i=0; i<num1.length; i++) {
        let mul = num1[num1.length-1-i] * num2 + flag
        res = String(mul % 10) + res
        flag = parseInt(mul / 10)
    }
    return flag ? String(flag) + res : res
}

var addStrings = function (num1, num2) {
  let a = num1
  let b = num2
  if (a.length > b.length) {
    b = '0'.repeat(a.length - b.length) + b
  } else {
    a = '0'.repeat(b.length - a.length) + a
  }
  let flag = 0
  let result = ''
  for (let i = 0; i < a.length; i++) {
    let sum = Number(a[a.length - 1 - i]) + Number(b[b.length - 1 - i]) + flag
    flag = sum >= 10 ? 1 : 0
    sum = sum % 10
    result = String(sum) + result
  }
  return flag === 1 ? '1' + result : result // 记得如果还有进位要加上
};