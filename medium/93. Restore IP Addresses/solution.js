/**
 * @param {string} s
 * @return {string[]}
 */
var restoreIpAddresses = function(s) {
    if(s.length > 12 || s.length < 4) return []
    const ips = []
    const findIps = (pos, sec, ip) => {
        if(sec === 4 && pos === s.length) {
            ips.push(ip)
            return
        }
        for(let i=0; i<3; i++) {
            if(pos + i + 1 > s.length) return
            let section = s.slice(pos, pos+i+1)
            if(Number(section) > 255 || (section.length > 1 && section.startsWith('0'))) return
            findIps(pos+i+1, sec+1, ip === '' ? `${section}` : `${ip}.${section}`)
        }
    }
    findIps(0, 0, '')
    return ips
};