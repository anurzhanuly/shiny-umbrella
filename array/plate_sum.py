def solution(l):
    max_sum = 0
    result = []

    for index, val in enumerate(l):
        tmp_result = [str(val)]
        tmp_sum = val

        for idx, v in enumerate(l):
            if idx == index:
                continue

            tmp_sum += v
            tmp_result.append(str(v))

            if tmp_sum % 3 == 0 and tmp_sum > max_sum:
                max_sum = tmp_sum
                result = tmp_result[:]

    result.sort(reverse=True)

    if len(result):
        return int(''.join(result))

    return 0


