package tokenizer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	_ TokenType = iota
	TokenLParen
	TokenRParen
	TokenSymbol
	TokenNumber
	TokenString
)

var (
	ErrSyntax = errors.New("Syntax error")
)

type Tokener interface {
	Type() TokenType
}

type LParenToken struct {}
type RParenToken struct {}
type SymbolToken struct {
	Value string
}
type NumberToken struct {
	Value int32
}
type StringToken struct {
	Value string
}

func (LParenToken) Type() TokenType  { return TokenLParen }
func (RParenToken) Type() TokenType  { return TokenRParen }
func (SymbolToken) Type() TokenType { return TokenSymbol }
func (NumberToken) Type() TokenType { return TokenNumber }
func (StringToken) Type() TokenType { return TokenString }

func isSpecialChar(ch byte) bool {
	return ch == '(' || ch == ')' || ch == '"'
}

func Tokenize(input string) ([]Tokener, error) {
	var tokens []Tokener

	var builder strings.Builder

	for i := 0; i < len(input); i++ {
		ch := input[i]

		switch {
		case unicode.IsSpace(rune(ch)):
			continue
		case ch == '(':
			tokens = append(tokens, LParenToken{})
		case ch == ')':
			tokens = append(tokens, RParenToken{})
		case ch == '"':
			builder.Reset()
			i++
			for i < len(input) && input[i] != '"' {
				builder.WriteByte(input[i])
				i++
			}
			tokens = append(tokens, StringToken{Value: builder.String()})
		case unicode.IsDigit(rune(ch)):
			builder.Reset()
			for i < len(input) && !unicode.IsSpace(rune(input[i])) && !isSpecialChar(input[i]) {
				if !unicode.IsDigit(rune(input[i])) {
					return nil, fmt.Errorf("%w: invalid character in number: %c", ErrSyntax, input[i])
				}
				builder.WriteByte(input[i])
				i++
			}
			i--
			numValue, err := strconv.Atoi(builder.String())
			if err != nil {
				return nil, fmt.Errorf("%w: invalid number: %s", ErrSyntax, builder.String())
			}
			tokens = append(tokens, NumberToken{Value: int32(numValue)})
		default:
			builder.Reset()
			for i < len(input) && !unicode.IsSpace(rune(input[i])) && !isSpecialChar(input[i]) {
				builder.WriteByte(input[i])
				i++
			}
			i--
			tokens = append(tokens, SymbolToken{Value: builder.String()})
		}
	}
	return tokens, nil
}