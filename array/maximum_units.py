class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        # 1: sort the botTypes
        result: int = 0
        max_value: int = 0
        for values in boxTypes:
            if max_value < values[1]:
                max_value = values[1]
        #  1.1: using counting sort
        count_array = [0] * (max_value+1)
        
        for values in boxTypes:
            count_array[values[1]] += values[0]
            
        for index in range(max_value):
            i = max_value-index
            
            if truckSize == 0:
                return result
            
            if count_array[i] > truckSize:
                return result + (truckSize * i)
        
            result += count_array[i] * i
            truckSize -= count_array[i]
            
        return result
                
            