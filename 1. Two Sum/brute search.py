class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for idx, e in enumerate(nums):
            searching = target - e
            for j in range(idx+1,len(nums)):
                if nums[j] == searching:
                    return [idx, j]