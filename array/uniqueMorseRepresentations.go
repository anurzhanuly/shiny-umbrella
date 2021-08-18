package array

func UniqueMorseRepresentations(words []string) int {
	var result int
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	encryptedWords := []string{}

	for index, word := range words {
		for _, letter := range word {
			encryptedWords[index] += morse[letter-'a']
		}
	}

	return result
}
