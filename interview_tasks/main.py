def get_test_cases(counter, test_cases):
    if counter == 0:
        return test_cases

    _ = input()
    test_values = input()
    test_cases.append(test_values.split())

    return get_test_cases(counter - 1, test_cases)


def square(values):
    if not len(values):
        return 0

    value = int(values.pop())

    return (value * value) + square(values)


def get_square_roots(test_cases, result):
    if not len(test_cases):
        return result

    single_case = test_cases.pop()
    result.append(square(single_case))

    return get_square_roots(test_cases, result)


def print_values(values):
    pass


def main():
    container = []
    cases_count = int(input())
    test_cases = get_test_cases(cases_count, container)
    squared_values = get_square_roots(test_cases, container)

    print(squared_values)
    # print_values(squared_values)


if __name__ == "__main__":
    main()
