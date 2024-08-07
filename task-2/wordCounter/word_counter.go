package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordCounter(words string) map[string]int{
	words = strings.ToLower(words)
	listWords := strings.Split(words, " ")
	var counter = make(map[string]int)


	for _, word := range listWords{
		var str string
		for _, ch := range word{
			if unicode.IsLetter(ch){
				str += string(ch)
			}
		}
		counter[str] += 1
	}
	return counter
}

func main(){

	// a simple example
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordCounter(text))	

	// this test is to simulate if it pass for different punctuation marks
	text2 := "The qu----%ick quick QUICK Qu-12-ICK"
	fmt.Println(wordCounter(text2))
}
