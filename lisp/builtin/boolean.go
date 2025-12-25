package builtin

import (
	"fmt"

	"github.com/Luke256/mylisp/lisp/value"
)

func BooleanAnd(args []value.Valuer) (value.Valuer, error) {
	for _, arg := range args {
		boolVal, ok := arg.(*value.Boolean)
		if !ok {
			return nil, fmt.Errorf("%w: expected Boolean, got %T", ErrTypeMismatch, arg)
		}
		if !boolVal.Value {
			return &value.Boolean{Value: false}, nil
		}
	}
	return &value.Boolean{Value: true}, nil
}

func BooleanOr(args []value.Valuer) (value.Valuer, error) {
	for _, arg := range args {
		boolVal, ok := arg.(*value.Boolean)
		if !ok {
			return nil, fmt.Errorf("%w: expected Boolean, got %T", ErrTypeMismatch, arg)
		}

		if boolVal.Value {
			return &value.Boolean{Value: true}, nil
		}
	}
	return &value.Boolean{Value: false}, nil
}

func BooleanNot(args []value.Valuer) (value.Valuer, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("not expects 1 argument, got %d", len(args))
	}
	boolVal, ok := args[0].(*value.Boolean)
	if !ok {
		return nil, fmt.Errorf("%w: expected Boolean, got %T", ErrTypeMismatch, args[0])
	}
	return &value.Boolean{Value: !boolVal.Value}, nil
}

func BooleanEqual(args []value.Valuer) (value.Valuer, error) {
	if len(args) < 2 {
		return &value.Boolean{Value: true}, nil
	}
	
	for _, arg := range args[1:] {
		if !args[0].Equal(arg) {
			return &value.Boolean{Value: false}, nil
		}
	}

	return &value.Boolean{Value: true}, nil
}