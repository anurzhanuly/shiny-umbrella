from typing import List


def count_sort(nums: List[int]) -> List[int]:
    max_val = max(nums)
    
    count_array = [0] * (max_val + 1)
    result = [0] * len(nums)
    
    for value in nums:
        count_array[value] += 1
        
    for index in range(1, len(count_array)):
        count_array[index] += count_array[index-1]
        
    for index, value in enumerate(nums):
        count_array[value] -= 1
        result[count_array[value]] = value
        
    return result
