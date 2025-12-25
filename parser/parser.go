package parser

import (
	"errors"
	"fmt"

	"github.com/Luke256/mylisp/tokenizer"
)

var (
	ErrUnexpectedEOF = errors.New("unexpected end of expression")
	ErrInvalidToken  = errors.New("invalid token")
)

func Parse(tokens []tokenizer.Tokener) ([]Exprer, error) {
	var exprs []Exprer

	var pos int = 0
	for pos < len(tokens) {
		expr, newPos, err := parseExpr(tokens, pos)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)
		pos = newPos
	}

	return exprs, nil
}

func parseExpr(tokens []tokenizer.Tokener, pos int) (Exprer, int, error) {
	switch tok := tokens[pos].(type) {
	case tokenizer.NumberToken:
		return &Number{Value: tok.Value}, pos + 1, nil
	case tokenizer.SymbolToken:
		return &Symbol{Name: tok.Value}, pos + 1, nil
	case tokenizer.LParenToken:
		pos++
		if pos >= len(tokens) {
			return nil, pos, ErrUnexpectedEOF
		}

		var args []Exprer
		for pos < len(tokens) {
			if _, ok := tokens[pos].(tokenizer.RParenToken); ok {
				pos++
				return &List{Exprs: args}, pos, nil
			}
			arg, newPos, err := parseExpr(tokens, pos)
			if err != nil {
				return nil, pos, err
			}
			args = append(args, arg)
			pos = newPos
		}
		return nil, pos, ErrUnexpectedEOF
	default:
		return nil, pos, fmt.Errorf("%w: %v (%v)", ErrInvalidToken, tok, tok.Type())
	}
}
