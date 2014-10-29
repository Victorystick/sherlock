package parser

import (
	"testing"
)

func match(t *testing.T, expected, tokens []Token) {
	if len(expected) != len(tokens) {
		t.Log(tokens)
		t.Fatalf("Didn't parse expected num (%d) of tokens, parsed (%d)",
			len(expected),
			len(tokens))
	}

	for i, expected := range expected {
		if expected.Type != tokens[i].Type || expected.Value != tokens[i].Value {
			t.Fatalf("Expected %v to have been %v", tokens[i], expected)
		}
	}
}

func TestLexStream(t *testing.T) {
	lexer := NewLexer()

	lexer.Stream([]byte("i 13 for { my -152"))

	lexer.Lex()

	tokens := lexer.Tokens()

	expectedTokens := []Token{
		Token{
			Type:  TokId,
			Value: "i",
		},
		Token{
			Type:  TokNum,
			Value: "13",
		},
		Token{
			Type:  TokId,
			Value: "for",
		},
		Token{
			Type:  TokLBrack,
			Value: "{",
		},
		Token{
			Type:  TokId,
			Value: "my",
		},
		Token{
			Type:  TokNum,
			Value: "-152",
		},
		Token{
			Type:  TokEOF,
			Value: "<EOF>",
		},
	}

	match(t, expectedTokens, tokens)
}

func TestGenericJava(t *testing.T) {
	tokens, err := Tokenize("class List<T> { T val; List<T> next; }")

	if err != nil {
		t.Fatal(err)
	}

	expectedTokens := []Token{
		Token{
			Type: TokId,
			Value: "class",
		},
		Token{
			Type: TokId,
			Value: "List",
		},
		Token{
			Type: TokPunc,
			Value: "<",
		},
		Token{
			Type: TokId,
			Value: "T",
		},
		Token{
			Type: TokPunc,
			Value: ">",
		},
		Token{
			Type: TokLBrack,
			Value: "{",
		},
		Token{
			Type: TokId,
			Value: "T",
		},
		Token{
			Type: TokId,
			Value: "val",
		},
		Token{
			Type: TokPunc,
			Value: ";",
		},
		Token{
			Type: TokId,
			Value: "List",
		},
		Token{
			Type: TokPunc,
			Value: "<",
		},
		Token{
			Type: TokId,
			Value: "T",
		},
		Token{
			Type: TokPunc,
			Value: ">",
		},
		Token{
			Type: TokId,
			Value: "next",
		},
		Token{
			Type: TokPunc,
			Value: ";",
		},
		Token{
			Type: TokRBrack,
			Value: "}",
		},
		Token{
			Type: TokEOF,
			Value: "<EOF>",
		},
	}

	match(t, expectedTokens, tokens)
}

func TestTokenize(t *testing.T) {
	tokens, _ := Tokenize("if (a > 7) { b = 7 }")

	expectedTokens := []Token{
		Token{
			Type:  TokId,
			Value: "if",
		},
		Token{
			Type:  TokPunc,
			Value: "(",
		},
		Token{
			Type:  TokId,
			Value: "a",
		},
		Token{
			Type:  TokPunc,
			Value: ">",
		},
		Token{
			Type:  TokNum,
			Value: "7",
		},
		Token{
			Type:  TokPunc,
			Value: ")",
		},
		Token{
			Type:  TokLBrack,
			Value: "{",
		},
		Token{
			Type:  TokId,
			Value: "b",
		},
		Token{
			Type:  TokPunc,
			Value: "=",
		},
		Token{
			Type:  TokNum,
			Value: "7",
		},
		Token{
			Type:  TokRBrack,
			Value: "}",
		},
		Token{
			Type:  TokEOF,
			Value: "<EOF>",
		},
	}

	match(t, expectedTokens, tokens)
}
