package palindrome_test

import (
	"palindrome/src/palindrome"
	"testing"
)

func TestIfWordArePalindrome(t *testing.T) {
	tests := []palindrome.Word{
		"Rotator",
		"bob",
		"madam",
		"mAlAyAlam",
		"1",
		"Able was I, ere I saw Elba",
		"Madam I’m Adam",
		"Step on no pets.",
		"Top spot!",
		"02/02/2020",
		"Socorram-me subi no ônibus em marrocos",
	}

	for _, test := range tests {
		want := true
		got := test.IsPalindrome()

		if want != got {
			t.Fatalf("Te current word %v are not a palindrome", test)
		}
	}
}

func TestIfWordAreNotAPalindrome(t *testing.T) {
	tests := []palindrome.Word{
		"xyz",
		"elephant",
		"Country",
		"Top . post!",
		"Wonderful-fool",
		"Wild imagination!",
	}

	for _, test := range tests {
		want := false
		got := test.IsPalindrome()

		if want != got {
			t.Fatalf("Te current word %v are a palindrome", test)
		}
	}
}

func TestReverseWord(t *testing.T) {
	test := palindrome.Word("123456789")
	want := palindrome.Word("987654321")

	if want != test.ReverseWord() {
		t.Fatalf("error ReverseWord")
	}

}

func TestToLowerWord(t *testing.T) {
	test := palindrome.Word("TEST")
	want := palindrome.Word("test")

	if test.ToLower() != want {
		t.Fatalf("error ToLowerWord")
	}
}

func TestRuleReplaceToBePalindrome(t *testing.T) {
	test := palindrome.Word("testúnitâ-ri/ô")
	want := palindrome.Word("testúnitâriô")

	if test.RuleReplaceToBePalindrome() != want {
		t.Fatalf("error RuleReplaceToBePalindrome")
	}
}

func TestRemoveAccents(t *testing.T) {
	test := palindrome.Word("testúnitâriô")
	want := palindrome.Word("testunitario")

	if test.RemoveAccents() != want {
		t.Fatalf("error RemoveAccents")
	}
}
