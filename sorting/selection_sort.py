from typing import List


def selection_sort(nums: List[int]) -> List[int]:
    for index in range(len(nums)):
        tmp = find_smallest_index(nums[index:]) + index

        nums[index], nums[tmp] = nums[tmp], nums[index]

    return nums


def find_smallest_index(nums: List[int]) -> int:
    tmp = 0

    for i, val in enumerate(nums):
        if val < nums[tmp]:
            tmp = i

    return tmp

nums = [2, 10, -1, 12, 111, 23, 5, 6]
print(selection_sort(nums))
