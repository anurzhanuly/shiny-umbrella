def get_test_cases(counter, test_cases):
    if counter == 0:
        return test_cases

    _ = input()
    test_values = input()
    test_cases.append(test_values)

    return get_test_cases(counter - 1, test_cases)


def get_square_roots(values):
    pass


def print_values(values):
    pass


def main():
    cases_count = int(input())
    test_cases = get_test_cases(cases_count, [])

    print(test_cases)
    # squared_values = get_square_roots(test_cases)
    # print_values(squared_values)


if __name__ == "__main__":
    main()
