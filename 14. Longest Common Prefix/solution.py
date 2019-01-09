class Solution:
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        length = len(strs)
        if(length == 0):
            return ''
        common = strs[0]    # init common prefix as the first str
        for i in range(1, length):  # compare the second and continue one by one, slice the common
            temp = ''
            for j in range(len(strs[i])):
                if j >= len(common):
                    break
                if strs[i][j] == common[j]:
                    temp += common[j]
                else:
                    break
            if(temp == ''):
                return temp
            common = temp
        return common