package builtin

import (
	"fmt"

	"github.com/Luke256/mylisp/lisp/value"
)

func Cons(args []value.Valuer) (value.Valuer, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("cons expects 2 arguments, got %d", len(args))
	}

	return &value.List{
		A: args[0],
		B: args[1],
	}, nil
}

func Car(args []value.Valuer) (value.Valuer, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("car expects 1 argument, got %d", len(args))
	}

	listVal, ok := args[0].(*value.List)
	if !ok {
		return nil, fmt.Errorf("car expects a list argument, got %T", args[0])
	}

	return listVal.A, nil
}

func Cdr(args []value.Valuer) (value.Valuer, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("cdr expects 1 argument, got %d", len(args))
	}
	listVal, ok := args[0].(*value.List)
	if !ok {
		return nil, fmt.Errorf("cdr expects a list argument, got %T", args[0])
	}
	return listVal.B, nil
}

func List(args []value.Valuer) (value.Valuer, error) {
	if len(args) == 0 {
		return &value.Unit{}, nil
	}
	head := &value.List{
		A: args[0],
	}
	current := head
	for _, arg := range args[1:] {
		newList := &value.List{
			A: arg,
		}
		current.B = newList
		current = newList
	}
	current.B = &value.Unit{}
	return head, nil
}

func NullP(args []value.Valuer) (value.Valuer, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("null? expects 1 argument, got %d", len(args))
	}
	_, isUnit := args[0].(*value.Unit)
	return &value.Boolean{Value: isUnit}, nil
}