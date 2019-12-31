package lexer

import (
	"fmt"
	"testing"
)

func TestNewSimpleLexer(t *testing.T) {
	l := NewSimpleLexer()
	err := l.Tokenize("age >= 45")
	if err != nil {
		t.Error(err)
	}
	for _, token := range l.tokens {
		fmt.Println(token.text, ":", TokenMap[token._type])
	}

	fmt.Println("------")

	l = NewSimpleLexer()
	err = l.Tokenize("int a = 1")
	if err != nil {
		t.Error(err)
	}
	for _, token := range l.tokens {
		fmt.Println(token.text, ":", TokenMap[token._type])
	}
}
