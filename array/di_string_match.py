import sys

class Solution:
    def diStringMatch(self, s: str) -> List[int]:
        max = sys.maxsize
        min = -max - 1
        result: list = []
        
        for char in s:
            if char == "I":
                result.append(min)
                min += 1
            else:
                result.append(max)
                max -= 1
        
        return result