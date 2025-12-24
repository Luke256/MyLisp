package lisp

import (
	"github.com/Luke256/mylisp/tokenizer"
	"github.com/Luke256/mylisp/parser"
	"github.com/Luke256/mylisp/expression"
)

type Box struct {
	symbols map[string]func(args []expression.Exprer) (expression.Exprer, error)
}

func NewBox() *Box {
	return &Box{
		symbols: make(map[string]func(args []expression.Exprer) (expression.Exprer, error)),
	}
}

func (b *Box) Eval(input string) error {
	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		return err
	}

	exprs, err := parser.Parse(tokens)
	if err != nil {
		return err
	}

	for _, expr := range exprs {
		_, err := b.evalExpr(expr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Box) evalExpr(expr expression.Exprer) (expression.Exprer, error) {
	return expr, nil
}