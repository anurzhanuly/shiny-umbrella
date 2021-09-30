def insertion_sort(nums: int):
    for j in range(1, len(nums)):
        i = j - 1
        current_number = nums[j]

        while i >= 0 and nums[i] > current_number:
            nums[i+1], nums[i] = nums[i], nums[i+1]
            i -= 1

    return nums

print(insertion_sort([31, 48, 56, 21, 41, 58]))
