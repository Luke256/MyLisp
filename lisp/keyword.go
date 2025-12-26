package lisp

import (
	"fmt"

	"github.com/Luke256/mylisp/lisp/value"
	"github.com/Luke256/mylisp/parser"
)

func (b *Box) registerKeywords() {
	b.vars["lambda"] = &value.KeyWord{Name: "lambda"}
	b.vars["define"] = &value.KeyWord{Name: "define"}
	b.vars["if"] = &value.KeyWord{Name: "if"}
	b.vars["let"] = &value.KeyWord{Name: "let"}
}

func (b *Box) evalKeyword(keyVal *value.KeyWord, args []parser.Exprer) (value.Valuer, error) {
	switch keyVal.Name {
	case "lambda":
		return b.keyLambda(args)
	case "define":
		return b.keyDefine(args)
	case "if":
		return b.keyIf(args)
	case "let":
		return b.keyLet(args)
	default:
		return nil, fmt.Errorf("unknown keyword: %s", keyVal.Name)
	}
}

// (lambda (args ...) body)
func (b *Box) keyLambda(args []parser.Exprer) (value.Valuer, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("lambda expects 2 arguments, got %d", len(args))
	}

	// parse argument list
	argListExpr, ok := args[0].(*parser.List)
	if !ok {
		return nil, fmt.Errorf("lambda argument list must be a list, got %q (%T)", args[0], args[0])
	}

	var argNames []string
	for _, argExpr := range argListExpr.Exprs {
		sym, ok := argExpr.(*parser.Ident)
		if !ok {
			return nil, fmt.Errorf("lambda argument names must be symbols, got %q (%T)", argExpr, argExpr)
		}
		argNames = append(argNames, sym.Name)
	}

	return &value.Function{
		Expr: args[1],
		Args: argNames,
	}, nil
}

// (define name value)
// or
// (define (func-name args ...) body) = (define func-name (lambda (args ...) body))
func (b *Box) keyDefine(args []parser.Exprer) (value.Valuer, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("define expects 2 arguments, got %d", len(args))
	}

	// function definition shorthand
	if listExpr, ok := args[0].(*parser.List); ok {
		if len(listExpr.Exprs) == 0 {
			return nil, fmt.Errorf("define function name missing")
		}

		funcName, ok := listExpr.Exprs[0].(*parser.Ident)
		if !ok {
			return nil, fmt.Errorf("define function name must be a symbol, got %q (%T)", listExpr.Exprs[0], listExpr.Exprs[0])
		}

		lambdaExpr := &parser.List{
			Exprs: []parser.Exprer{
				&parser.Ident{Name: "lambda"},
				&parser.List{Exprs: listExpr.Exprs[1:]},
				args[1],
			},
		}

		args[0] = funcName
		args[1] = lambdaExpr
	}

	nameSym, ok := args[0].(*parser.Ident)
	if !ok {
		return nil, fmt.Errorf("define first argument must be a symbol, got %q (%T)", args[0], args[0])
	}

	val, err := b.evalExpr(args[1])
	if err != nil {
		return nil, err
	}

	b.Register(nameSym.Name, val)

	return &value.Unit{}, nil
}

// (if condition then else)
func (b *Box) keyIf(args []parser.Exprer) (value.Valuer, error) {
	if len(args) != 2 && len(args) != 3 {
		return nil, fmt.Errorf("if expects 2 or 3 arguments, got %d", len(args))
	}

	condVal, err := b.evalExpr(args[0])
	if err != nil {
		return nil, err
	}
	condBool, ok := condVal.(*value.Boolean)
	if !ok {
		return nil, fmt.Errorf("if condition must be a boolean, got %q (%T)", condVal, condVal)
	}

	if condBool.Value {
		return b.evalExpr(args[1])
	} else if len(args) == 3 {
		return b.evalExpr(args[2])
	}
	return &value.Unit{}, nil
}

// (let ((var1 val1) (var2 val2) ...) body1 body2 ...)
func (b *Box) keyLet(args []parser.Exprer) (value.Valuer, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("let expects at least 1 argument, got %d", len(args))
	}

	box := newChildBox(b)

	bindings, ok := args[0].(*parser.List)
	if !ok {
		return nil, fmt.Errorf("let first argument must be a list of bindings, got %q (%T)", args[0], args[0])
	}

	for _, bindExpr := range bindings.Exprs {
		bindPair, ok := bindExpr.(*parser.List)
		if !ok || len(bindPair.Exprs) != 2 {
			return nil, fmt.Errorf("let bindings must be lists of 2 elements, got %q (%T)", bindExpr, bindExpr)
		}

		varName, ok := bindPair.Exprs[0].(*parser.Ident)
		if !ok {
			return nil, fmt.Errorf("let binding name must be a symbol, got %q (%T)", bindPair.Exprs[0], bindPair.Exprs[0])
		}

		evaledVal, err := b.evalExpr(bindPair.Exprs[1])
		if err != nil {
			return nil, err
		}

		box.Register(varName.Name, evaledVal)
	}

	var result value.Valuer = &value.Unit{}
	var err error

	for _, bodyExpr := range args[1:] {
		result, err = box.evalExpr(bodyExpr)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
