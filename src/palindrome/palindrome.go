package palindrome

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type IWord interface {
	IsPalindrome() bool
	ReverseWord() Word
	ToLower() Word
	RuleReplaceToBePalindrome() Word
	RemoveAccents() Word
}

type Word string

func (w Word) IsPalindrome() bool {

	return w.ToLower().RuleReplaceToBePalindrome().RemoveAccents() == w.ReverseWord().ToLower().RuleReplaceToBePalindrome().RemoveAccents()

}

func (w Word) ReverseWord() (reverseWord Word) {
	for _, v := range w {
		reverseWord = Word(v) + reverseWord
	}
	return
}

func (w Word) ToLower() Word {
	return Word(strings.ToLower(string(w)))
}

func (w Word) RuleReplaceToBePalindrome() Word {
	convWord := string(w)

	convWord = strings.ReplaceAll(convWord, ",", "")
	convWord = strings.ReplaceAll(convWord, " ", "")
	convWord = strings.ReplaceAll(convWord, "’", "")
	convWord = strings.ReplaceAll(convWord, "/", "")
	convWord = strings.ReplaceAll(convWord, "-", "")
	convWord = strings.ReplaceAll(convWord, "^", "")
	convWord = strings.ReplaceAll(convWord, ".", "")
	convWord = strings.ReplaceAll(convWord, "!", "")

	return Word(convWord)
}

func (w Word) RemoveAccents() Word {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, string(w))
	if e != nil {
		panic(e)
	}
	return Word(output)
}
