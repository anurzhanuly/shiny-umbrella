class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        unique: dict = {}
        result: int = 0
            
        for value in nums:
            if value in unique:
                unique[value] = 0
            else:
                unique[value] = value
        
        for value in unique.values():
            result += value
            
        return result