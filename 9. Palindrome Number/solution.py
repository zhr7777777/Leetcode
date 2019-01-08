class Solution:
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0 or ( x % 10 == 0 and x != 0):      # filter nums like 10000
            return False
        result = 0
        origin = x
        while(x > result):  # revert half of x
            c = x % 10
            x = int(x / 10)
            result = result * 10 + c
        return x == result or x == int(result / 10) # if num of x is odd, like 12321, then x = 12, result = 123