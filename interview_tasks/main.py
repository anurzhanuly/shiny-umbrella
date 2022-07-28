def get_test_cases(counter: int, test_cases: list) -> list:
    if counter == 0:
        return test_cases

    _ = input()
    test_values = input()
    test_cases.append(test_values.split())

    return get_test_cases(counter - 1, test_cases)


def square(values: list) -> int:
    if not len(values):
        return 0

    value = int(values.pop())

    return (value * value) + square(values)


def get_square_roots(test_cases: list, result: list) -> list:
    if not len(test_cases):
        return result

    single_case = test_cases.pop()
    result.append(square(single_case))

    return get_square_roots(test_cases, result)


def print_values(values):
    if not len(values):
        return

    print(values.pop())

    print_values(values)


def main():
    cases_count = int(input())
    test_cases = get_test_cases(cases_count, [])
    squared_values = get_square_roots(test_cases, [])

    print(squared_values)
    print_values(squared_values)


if __name__ == "__main__":
    main()
