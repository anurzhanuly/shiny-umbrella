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
	// 1: convert "allowed" -> map
	allowedRunes := make(map[rune]int)
	letterCount := len(allowed)
	result := 0

	for _, rune := range allowed {
		allowedRunes[rune] = 0
	}

	// 2: iterate over the "words" array
	for i := 0; i < len(words); i++ {
		for _, letter := range words[i] {
			// 3: check if the letter of the array is in the map
			_, exists := allowedRunes[letter]

			if exists {
				allowedRunes[letter] = 1
			}
		}

		tmpCount := 0

		for key, value := range allowedRunes {
			tmpCount += value
			allowedRunes[key] = 0
		}

		if tmpCount == letterCount {
			result += 1
		}
	}

	return result
}
