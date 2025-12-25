package lisp

import (
	"errors"
	"fmt"

	"github.com/Luke256/mylisp/lisp/builtin"
	"github.com/Luke256/mylisp/lisp/value"
	"github.com/Luke256/mylisp/parser"
	"github.com/Luke256/mylisp/tokenizer"
)

var (
	ErrUndefinedSymbol = errors.New("undefined symbol")
)

type Box struct {
	vars   map[string]value.Valuer
	parent *Box
}

func NewBox() *Box {
	box := &Box{
		vars:   make(map[string]value.Valuer),
		parent: nil,
	}

	// add keyword
	box.registerKeywords()

	// add build-in functions
	box.AddBuildIn("+", builtin.ArithmeticAdd)
	box.AddBuildIn("-", builtin.ArithmeticSubtract)
	box.AddBuildIn("*", builtin.ArithmeticMultiply)
	box.AddBuildIn("/", builtin.ArithmeticDivide)

	return box
}

func newChildBox(parent *Box) *Box {
	return &Box{
		vars:   make(map[string]value.Valuer),
		parent: parent,
	}
}

func (b *Box) Register(name string, value value.Valuer) {
	b.vars[name] = value
}

func (b *Box) Resolve(name string) (value.Valuer, bool) {
	val, ok := b.vars[name]
	if ok {
		return val, true
	}
	if b.parent != nil {
		return b.parent.Resolve(name)
	}
	return nil, false
}

func (b *Box) AddBuildIn(name string, function func(args []value.Valuer) (value.Valuer, error)) {
	b.Register(name, &value.BuiltinFunction{
		Func: function,
	})
}

func (b *Box) Eval(input string) (value.Valuer, error) {
	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		return nil, err
	}

	exprs, err := parser.Parse(tokens)
	if err != nil {
		return nil, err
	}

	var val value.Valuer = &value.Unit{}
	for _, expr := range exprs {
		val, err = b.evalExpr(expr)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (b *Box) evalExpr(expr parser.Exprer) (value.Valuer, error) {
	switch e := expr.(type) {
	case *parser.Number:
		return &value.Number{Value: e.Value}, nil
	case *parser.String:
		return &value.String{Value: e.Value}, nil
	case *parser.Symbol:
		val, ok := b.Resolve(e.Name)
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrUndefinedSymbol, e.Name)
		}
		return val, nil
	case *parser.List:
		return b.evalCall(e)
	default:
		return nil, fmt.Errorf("unknown expression type: %T", expr)
	}
}

func (b *Box) evalCall(list *parser.List) (value.Valuer, error) {
	fVal, err := b.evalExpr(list.Exprs[0])
	if err != nil {
		return nil, err
	}

	switch fVal := fVal.(type) {
	case *value.Function:
		return b.evalFunction(fVal, list.Exprs[1:])
	case *value.BuiltinFunction:
		return b.evalBuildInFunction(fVal, list.Exprs[1:])
	case *value.KeyWord:
		return b.evalKeyword(fVal, list.Exprs[1:])
	default:
		return nil, fmt.Errorf("called value is not a function or keyword: %T", fVal)
	}
}

func (b *Box) evalBuildInFunction(builtinFunc *value.BuiltinFunction, argExprs []parser.Exprer) (value.Valuer, error) {
	var argVals []value.Valuer
	for _, argExpr := range argExprs {
		argVal, err := b.evalExpr(argExpr)
		if err != nil {
			return nil, err
		}
		argVals = append(argVals, argVal)
	}
	return builtinFunc.Func(argVals)
}

func (b *Box) evalFunction(function *value.Function, argExprs []parser.Exprer) (value.Valuer, error) {
	if len(function.Args) != len(argExprs) {
		return nil, fmt.Errorf("argument count mismatch: expected %d, got %d", len(function.Args), len(argExprs))
	}

	newBox := newChildBox(b)

	for i, argName := range function.Args {
		argVal, err := b.evalExpr(argExprs[i])
		if err != nil {
			return nil, err
		}
		newBox.Register(argName, argVal)
	}

	return newBox.evalExpr(function.Expr)
}
