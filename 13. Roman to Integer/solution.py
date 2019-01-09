class Solution:
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        mapping = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }
        remaining = s
        result = 0
        for i in range(len(s)):
            if i + 1 < len(s):
                if mapping[s[i + 1]] > mapping[s[i]]:   # if next > cur, result - cur, else result + cur
                    result -= mapping[s[i]]
                else:
                    result += mapping[s[i]]
            else:
                result += mapping[s[i]]
        return result