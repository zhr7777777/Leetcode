class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        h = {}  # 字典记录遍历过的元素，key为元素，value为索引值
        for idx, e in enumerate(nums):  # 只需遍历一次
            if target - e in h: # 从遍历过的元素找
                return [h[target - e], idx]
            h[e] = idx