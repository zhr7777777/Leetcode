/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var checkInclusion = function (s1, s2) {
    let map = {}
    for (let i = 0; i < s1.length; i++) {
      map[s1[i]] = s1[i] in map ? map[s1[i]] + 1 : 1
    }
    for (let i = 0; i < s2.length; i++) {
      if (!(s2[i] in map)) continue
      if (i + s1.length > s2.length) break
      let candidate = s2.slice(i, i + s1.length)
      if (isPermutation(map, candidate)) {
        return true
      }
    }
    return false
  };
  
  const isPermutation = function (map, s2) {
    const localMap = { ...map }
    for (let i = 0; i < s2.length; i++) {
      if (!(s2[i] in localMap)) return false
      if (localMap[s2[i]] === 0) return false
      localMap[s2[i]] = localMap[s2[i]] - 1
    }
    return true
  }