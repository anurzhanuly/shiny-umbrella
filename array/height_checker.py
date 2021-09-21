from typing import List


class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        counter: int = 0
        expected = sorted(heights)
        
        for index in range(len(heights)):
            if expected[index] != heights[index]:
                counter += 1
                
        return counter

    def heightCheckerOpt(self, heights: List[int]) -> int:
        # non negative values in the array

        # 1. find max value in array
        max_value: int = 0

        for value in heights:
            if value > max_value:
                max_value = value

        max_value += 1

        # 2. construct an array counter for each value
        # indexes represents value of the original array
        # each time value is repeated the value of corresponding index in 
        # counter array increments by 1
        counter_array = [0] * max_value

        for value in heights:
            counter_array[value] += 1
        
        # 3. add values in "counter" array such that
        # counter[i] = counter[i] + counter[i-1]
        for index in range(1, max_value):
            counter_array[index] = counter_array[index] + counter_array[index-1]
        
        # 4. based on that values of counter, iterate through original array
        # and place corresponding index's value of "counter" array
        # to newly created array with same length as the "origin" array
        sorted_values = [0] * len(heights)
        
        for value in heights:
            counter_array[value] -= 1
            sorted_values[counter_array[value]] = value
        # Result: array is sorted.

        return sorted_values