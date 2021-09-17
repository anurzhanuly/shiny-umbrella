def epi_start(book_titles: list, citation_counts: list) -> int:
    answer: dict = {}
    ans: int = 0

    for number in range(1, len(citation_counts)+1):
        for value in citation_counts:
            
            if number <= value:
                if number in answer:
                    answer[number] += 1
                else:
                    answer[number] = 1
    print(answer)

    for key in answer:
        if key > ans and answer[key] >= key:
            ans = key

    return ans

print(epi_start(book_titles=['a','s','b','c','d'], citation_counts=[1, 3, 4, 2, 6]))
print(epi_start(book_titles=['a','s','b','c','d','s','b','c','d'], citation_counts=[1, 4, 1, 4, 3, 2, 1, 5, 6]))