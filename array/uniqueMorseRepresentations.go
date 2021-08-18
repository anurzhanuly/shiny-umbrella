package array

import "strings"

func UniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	encryptedWords := make(map[string]bool)
	var morseWord strings.Builder

	for _, word := range words {
		morseWord.Reset()

		for _, letter := range word {
			morseWord.WriteString(morse[letter-'a'])
		}

		encryptedWords[morseWord.String()] = true
	}

	return len(encryptedWords)
}
