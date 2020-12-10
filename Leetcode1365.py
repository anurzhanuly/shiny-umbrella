def smallerNumbersThanCurrent(self, nums):
    nums_compared = [0] * len(nums)
    for i in range(len(nums)):
        for j in range(len(nums)):
            if i == j:
                continue
            if nums[i] > nums[j]:
                nums_compared[i] += 1
    return nums_compared
