from typing import List


class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        nums = sorted(nums)
        result: int = 0
        
        for index in range(len(nums)):
            if index % 2 == 0:
                result += nums[index]
                
        return result