package builtin

import (
	"fmt"

	"github.com/Luke256/mylisp/lisp/value"
)

func Expt(args []value.Valuer) (value.Valuer, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expt expects 2 arguments, got %d", len(args))
	}
	baseVal, ok := args[0].(*value.Number)
	if !ok {
		return nil, fmt.Errorf("expt expects number arguments, got %T", args[0])
	}
	expVal, ok := args[1].(*value.Number)
	if !ok {
		return nil, fmt.Errorf("expt expects number arguments, got %T", args[1])
	}
	if expVal.Value < 0 {
		return nil, fmt.Errorf("expt does not support negative exponents")
	}
	result := int32(1)
	for i := int32(0); i < expVal.Value; i++ {
		result *= baseVal.Value
	}
	return &value.Number{Value: result}, nil
}