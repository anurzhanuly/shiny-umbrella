package array

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	var letter1, letter2, len1, len2 int

	for len(word1) > len1 && len(word2) > len2 {
		if word1[len1][letter1] != word2[len2][letter2] {
			return false
		}

		if len(word1[len1])-1 == letter1 {
			len1 += 1
			letter1 = 0
		} else {
			letter1 += 1
		}

		if len(word2[len2])-1 == letter2 {
			len2 += 1
			letter2 = 0
		} else {
			letter2 += 1
		}
	}

	if len(word1) > len1 || len(word2) > len2 {
		return false
	}

	return true
}

func CountConsistentStrings(allowed string, words []string) int {
	allowedRunes := make(map[rune]int)
	result := 0

	for _, rune := range allowed {
		allowedRunes[rune] = 0
	}

	for i := 0; i < len(words); i++ {

		isConsistent := true

		for _, letter := range words[i] {
			_, exists := allowedRunes[letter]

			if !exists {
				isConsistent = false

				break
			}
		}

		if isConsistent {
			result += 1
		}

	}

	return result
}
