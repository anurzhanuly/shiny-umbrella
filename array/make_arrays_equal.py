from typing import List


class Solution:
    def canBeEqual(self, target: List[int], arr: List[int]) -> bool:
        nums: dict = {}
            
        for value in target:
            if value in nums:
                nums[value] += 1
            else: 
                nums[value] = 1
        
        for value in arr:
            if value not in nums:
                return False
            
            nums[value] -= 1
            if nums[value] < 0:
                return False
            
        return True