class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        result = 0
        y = x
        if x < 0:
            y = abs(x)
        while(y):
            remainder = y % 10
            y = int(y / 10)
            result = result * 10 + remainder
            test = result if x > 0 else -result
            if test > pow(2, 31) - 1 or test < -pow(2, 31):
                return 0
        result = result if x > 0 else -result
        return result