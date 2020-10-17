package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func hangman(guessWord string, badguess int) bool {
	count := 0
	stringBuilder := make(map[rune]bool)
	guessWordRune := []rune(guessWord)

	for count < badguess && len(stringBuilder) < len(guessWord) {
		fmt.Println("Please enter a letter")
		in := bufio.NewReader(os.Stdin)
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		// isLetterBool判断读入的是否符合标准
		isLetterBool := unicode.IsLetter(r) && n == 1
		if err == io.EOF || !isLetterBool {
			fmt.Println("Error, please enter a letter.")
			continue
		}
		if isLetterBool {
			// 判断输入的字母是否在guessWordRune之中
			for _, value := range guessWordRune {
				if value == r {
					stringBuilder[r] = true
					fmt.Println("Congratulations, you guess the right letter!!!")
					break
				}
			}
		}
		count++
	}

	if len(stringBuilder) == len(guessWord) {
		return true
	} else {
		return false
	}
}

func main() {
	guessWord := "abcde"
	badguess := 7
	fmt.Println(hangman(guessWord, badguess))
}
