def smallerNumbersThanCurrent(nums):
        dicton = {}
        nums1 = sorted(nums, reverse=True)
        for index in range(0, len(nums1) - 1):
            if nums1[index] == nums1[index+1]:
                continue
            dicton[nums1[index]] = len(nums1) - index - 1
        dicton[nums1[-1]] = 0
        return [dicton[index] for index in nums]
