package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Function to check if a character is alphanumeric
func isAlphanumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsDigit(char)
}

// check if the string is palindrom
func palindrom(words string) bool{

	words = strings.ToLower(words) // making words to to be case insensitive
	listWords := strings.Split(words, " ")

	// creating a new string to check palindrom by removing non alphanumeric characters
	var str string  
	for _, word := range listWords{
		for _, ch := range word{
			if isAlphanumeric(ch){
				str += string(ch)
			}
		}
	}

	//checking palindrom by using two pointers technique
	left, right := 0, len(str) - 1
	for left < right{
		if str[left] != str[right]{
			return false
		}

		left += 1
		right -= 1
	}
	return true
}

func main(){
	fmt.Println(palindrom("A man, a plan, a canal, Panama!---"))
	fmt.Println(palindrom("honna"))
}