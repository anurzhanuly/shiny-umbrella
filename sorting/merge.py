def divide(arr):
    if len(arr) == 1:
        return arr

    middle_index = len(arr) // 2
    left = arr[0:middle_index]
    right = arr[middle_index:]

    left = divide(left)
    right = divide(right)

    return merge(left, right)


def merge(left, right):
    idx = 0
    idx_2 = 0
    res_idx = 0
    result = [0] * (len(left) + len(right))

    while idx < len(left) and idx_2 < len(right):
        left_val = left[idx]
        right_val = right[idx_2]

        if left_val < right_val:
            result[res_idx] = left_val
            idx += 1
        else:
            result[res_idx] = right_val
            idx_2 += 1

        res_idx += 1

    while idx < len(left):
        result[res_idx] = left[idx]
        idx += 1
        res_idx += 1

    while idx_2 < len(right):
        result[res_idx] = right[idx_2]
        idx_2 += 1
        res_idx += 1

    return result


def merge_sort(arr):
    return divide(arr)


print(merge_sort([10, 2, 120, 1111, -1, -212, 120, 0, 5]))
