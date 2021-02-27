class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        counter = 0
        def bin(row):
            start = 0
            end = len(row)-1
            
            while start < end:
                target_index = (start + end) // 2
                value = row[target_index]
                
                if value < 0:
                    end = target_index
                else:
                    start = target_index + 1
            if len(row):
                if row[start] < 0:
                    return len(row[start:])
            return 0
        
        for row in grid:
            counter += bin(row)
        
        return counter
