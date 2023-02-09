package main

import (
	"fmt"
	"palindrome/src/palindrome"
)

func main() {

	word := palindrome.Word("Socorram-me subi no ônibus em marrocos")

	fmt.Println(word.IsPalindrome())
}
