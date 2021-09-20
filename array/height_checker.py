class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        counter: int = 0
        expected = sorted(heights)
        
        for index in range(len(heights)):
            if expected[index] != heights[index]:
                counter += 1
                
        return counter