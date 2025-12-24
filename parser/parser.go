package parser

import (
	"errors"
	"fmt"

	"github.com/Luke256/mylisp/expression"
	"github.com/Luke256/mylisp/tokenizer"
)

var (
	ErrSyntax = errors.New("Syntax error")
)

func Parse(tokens []tokenizer.Tokener) ([]expression.Exprer, error) {
	var exprs []expression.Exprer
	var i int = 0

	for i < len(tokens) {
		expr, nextIndex, err := parseExpr(tokens, i)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)
		i = nextIndex
	}

	return exprs, nil
}

func parseExpr(tokens []tokenizer.Tokener, i int) (expression.Exprer, int, error) {
	if i >= len(tokens) {
		return nil, i, fmt.Errorf("%w: unexpected end of input", ErrSyntax)
	}

	token := tokens[i]

	switch t := token.(type) {
	case tokenizer.NumberToken:
		return &expression.NumberExpr{Value: t.Value}, i + 1, nil
	case tokenizer.StringToken:
		return &expression.StringExpr{Value: t.Value}, i + 1, nil
	case tokenizer.SymbolToken:
		return &expression.SymbolExpr{Name: t.Value}, i + 1, nil
	case tokenizer.LParenToken:
		var elements []expression.Exprer
		i++
		for i < len(tokens) && tokens[i].Type() != tokenizer.TokenRParen {
			elem, nextIndex, err := parseExpr(tokens, i)
			if err != nil {
				return nil, i, err
			}
			elements = append(elements, elem)
			i = nextIndex
		}

		if i >= len(tokens) || tokens[i].Type() != tokenizer.TokenRParen {
			return nil, i, fmt.Errorf("%w: expected ')'", ErrSyntax)
		}
		return &expression.ListExpr{Elements: elements}, i + 1, nil
	default:
		return nil, i, fmt.Errorf("%w: unexpected token", ErrSyntax)
	}
}
