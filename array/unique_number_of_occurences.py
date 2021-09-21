from collections import defaultdict
from typing import List

class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        occurences = defaultdict(int)
        unique = defaultdict(int)
        
        for number in arr:
            occurences[number] += 1
        
        for key, value in occurences.items():
            if unique[value] != 0:
                return False
            
            unique[value] = key
            
        return True