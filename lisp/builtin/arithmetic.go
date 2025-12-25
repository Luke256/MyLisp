package builtin

import (
	"fmt"

	"github.com/Luke256/mylisp/lisp/value"
)

func ArithmeticAdd(args []value.Valuer) (value.Valuer, error) {
	var sum int32 = 0
	for _, arg := range args {
		num, ok := arg.(*value.Number)
		if !ok {
			return nil, fmt.Errorf("%w: expected Number, got %T", ErrTypeMismatch, arg)
		}
		sum += num.Value
	}
	return &value.Number{Value: sum}, nil
}

func ArithmeticSubtract(args []value.Valuer) (value.Valuer, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("subtract expects at least one argument")
	}
	firstNum, ok := args[0].(*value.Number)
	if !ok {
		return nil, fmt.Errorf("%w: expected Number, got %T", ErrTypeMismatch, args[0])
	}
	result := firstNum.Value
	for _, arg := range args[1:] {
		num, ok := arg.(*value.Number)
		if !ok {
			return nil, fmt.Errorf("%w: expected Number, got %T", ErrTypeMismatch, arg)
		}
		result -= num.Value
	}
	return &value.Number{Value: result}, nil
}

func ArithmeticMultiply(args []value.Valuer) (value.Valuer, error) {
	result := int32(1)
	for _, arg := range args {
		num, ok := arg.(*value.Number)
		if !ok {
			return nil, fmt.Errorf("%w: expected Number, got %T", ErrTypeMismatch, arg)
		}
		result *= num.Value
	}
	return &value.Number{Value: result}, nil
}

func ArithmeticDivide(args []value.Valuer) (value.Valuer, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("divide expects at least one argument")
	}
	firstNum, ok := args[0].(*value.Number)
	if !ok {
		return nil, fmt.Errorf("%w: expected Number, got %T", ErrTypeMismatch, args[0])
	}
	result := firstNum.Value
	for _, arg := range args[1:] {
		num, ok := arg.(*value.Number)
		if !ok {
			return nil, fmt.Errorf("%w: expected Number, got %T", ErrTypeMismatch, arg)
		}
		if num.Value == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		result /= num.Value
	}
	return &value.Number{Value: result}, nil
}