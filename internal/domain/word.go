package hangman

import "fmt"

type Word struct {
	Word                []rune
	Difficulty          string
	Category            string
	CountGuessedLetters int8
	GuessedLetters      []bool
}

func (*Word) NewWord(word string, level string, category string) Word {
	return Word{Word: []rune(word),
		Difficulty:          level,
		Category:            category,
		CountGuessedLetters: 0,
		GuessedLetters:      make([]bool, len([]rune(word)))}
}

func (word *Word) UpdateGuessedLetters(char rune) {
	for index, letter := range word.Word {
		if letter == char {
			word.GuessedLetters[index] = true
			word.CountGuessedLetters++
		}

		if word.GuessedLetters[index] {
			fmt.Print(string(letter))
		} else {
			fmt.Print("_")
		}
	}
}
