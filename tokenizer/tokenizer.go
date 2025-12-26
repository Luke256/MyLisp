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
	TokenIdent
	TokenNumber
	TokenString
	TokenBoolean
)

var (
	ErrSyntax = errors.New("Syntax error")
)

type Tokener interface {
	Type() TokenType
}

type LParenToken struct{}
type RParenToken struct{}
type ItentToken struct {
	Value string
}
type NumberToken struct {
	Value int32
}
type StringToken struct {
	Value string
}
type BooleanToken struct {
	Value bool
}

func (LParenToken) Type() TokenType  { return TokenLParen }
func (RParenToken) Type() TokenType  { return TokenRParen }
func (ItentToken) Type() TokenType   { return TokenIdent }
func (NumberToken) Type() TokenType  { return TokenNumber }
func (StringToken) Type() TokenType  { return TokenString }
func (BooleanToken) Type() TokenType { return TokenBoolean }

func isSpecialChar(ch byte) bool {
	return ch == '(' || ch == ')' || ch == '"' || ch == '#'
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
		case ch == '#':
			i++
			if i >= len(input) {
				return nil, fmt.Errorf("%w: unexpected end after #", ErrSyntax)
			}

			switch input[i] {
			case 't':
				tokens = append(tokens, BooleanToken{Value: true})
			case 'f':
				tokens = append(tokens, BooleanToken{Value: false})
			default:
				return nil, fmt.Errorf("%w: invalid boolean literal: #%c", ErrSyntax, input[i])
			}
		default:
			builder.Reset()
			for i < len(input) && !unicode.IsSpace(rune(input[i])) && !isSpecialChar(input[i]) {
				builder.WriteByte(input[i])
				i++
			}
			i--
			tokens = append(tokens, ItentToken{Value: builder.String()})
		}
	}
	return tokens, nil
}
